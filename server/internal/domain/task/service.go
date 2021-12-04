package task

import (
	"context"
	"fmt"
	"math"
	"server/internal/config"
	"server/internal/domain"
	"server/internal/ent"
	"server/internal/ent/predicate"
	taskent "server/internal/ent/task"
	userent "server/internal/ent/user"
	"server/internal/platform"
	"sort"
)

type Service struct {
	client    *ent.TaskClient
	jwtConfig config.Jwt
}

func NewService(client *ent.Client, config config.Config) *Service {
	return &Service{
		client:    client.Task,
		jwtConfig: config.Jwt,
	}
}

func (s Service) ByID(ctx context.Context, taskID int) (*domain.Task, error) {
	task, err := s.client.Query().
		Where(taskent.ID(taskID)).
		WithCreator().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, platform.NotFound("Task is not found")
		}
		return nil, platform.WrapInternal(err)
	}

	return domain.TaskFromEnt(task), nil
}

func (s Service) Fetch(ctx context.Context, dto domain.GetTaskDTO) ([]*domain.Task, error) {
	query := s.client.Query().Select()
	if dto.Estimated != nil {
		query.Where(taskent.Estimated(*dto.Estimated))
	}
	if dto.Urgency != nil {
		query.Where(taskent.Urgency(int8(*dto.Urgency)))
	}
	if dto.Importance != nil {
		query.Where(taskent.Importance(int8(*dto.Importance)))
	}
	if dto.OrderBy != nil && dto.Order != nil {
		if *dto.Order == "asc" {
			query.Order(ent.Asc(*dto.OrderBy))
		} else {
			query.Order(ent.Desc(*dto.OrderBy))
		}
	}
	query.Where(taskent.HasCreatorWith(userent.ID(dto.UserID)))
	taskents, err := query.All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, platform.NotFound("Tasks are not found")
		}
		return nil, platform.WrapInternal(err)
	}
	if dto.Priority != nil {
		if !*dto.Priority {
			return domain.TasksFromEnt(taskents), nil
		}
		tasks := (domain.Tasks)(domain.TasksFromEnt(taskents))
		sort.Sort(sort.Reverse(tasks))
		return tasks, nil
	}
	return domain.TasksFromEnt(taskents), nil
}

func (s Service) Update(ctx context.Context, taskDTO domain.TaskPutDTO) (*domain.Task, error) {
	err := s.client.Update().
		Where(
			taskent.ID(taskDTO.TaskID),
			taskent.HasCreatorWith(userent.ID(taskDTO.UserID)),
		).
		SetNillableIcon(taskDTO.Icon).
		SetTitle(taskDTO.Title).
		SetNillableDescription(taskDTO.Description).
		SetNillableDeadline(taskDTO.Deadline).
		SetNillableEstimated(taskDTO.Estimated).
		SetUrgency(int8(taskDTO.Urgency)).
		SetImportance(int8(taskDTO.Importance)).
		Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, platform.NotFound("Task is not found")
		}
		return nil, platform.WrapInternal(err)
	}

	task, err := s.ByID(ctx, taskDTO.TaskID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s Service) Create(ctx context.Context, taskDTO domain.CreateTaskDTO) (*domain.Task, *domain.Question, error) {
	entTask, err := s.client.Create().
		SetTitle(taskDTO.Title).
		SetNillableIcon(taskDTO.Icon).
		SetNillableDescription(taskDTO.Description).
		SetImportance(int8(taskDTO.Importance)).
		SetUrgency(int8(taskDTO.Urgency)).
		SetNillableDeadline(taskDTO.Deadline).
		SetNillableEstimated(taskDTO.Estimated).
		SetCreatorID(taskDTO.UserID).
		SetLo(0).
		SetHi(0).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, nil, platform.NotFound("User not found")
		}
		return nil, nil, platform.WrapInternal(err)
	}
	creator, err := entTask.QueryCreator().Only(ctx)
	if err != nil {
		return nil, nil, platform.WrapInternal(err)
	}
	entTask.Edges.Creator = creator
	task := domain.TaskFromEnt(entTask)

	task.CustomMult = 1

	question, err := s.GenerateQuestion(ctx, task)
	if err != nil {
		return nil, nil, err
	}

	return task, question, nil
}

func (s Service) AnswerQuestion(ctx context.Context, params domain.AnswerQuestionDTO) (*domain.Question, error) {
	task1, err := s.client.Query().
		Where(taskent.ID(params.TaskID)).
		WithCreator().
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, platform.NotFound("Task is not found")
		}
		return nil, platform.WrapInternal(err)
	}
	task2, err := s.client.Query().
		Where(taskent.ID(params.CompareTaskID)).
		WithCreator().
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, platform.NotFound("Task is not found")
		}
		return nil, platform.WrapInternal(err)
	}

	var constant float64

	if task1.Hi == 0 || task1.Lo == 0 {
		switch params.Response {
		case 1:
			constant += Epsilon
			task1.Lo = task2.CustomMult
			newHiTask, err := s.client.Query().
				Where(taskent.CreatorID(params.UserID)).
				WithCreator().
				Order(ent.Desc()).
				First(ctx)
			if err != nil {
				return nil, platform.NotFound("Task is not found")
			}
			task1.Hi = newHiTask.CustomMult
			break
		case -1:
			constant -= Epsilon
			task1.Hi = task2.CustomMult
			newLoTask, err := s.client.Query().
				Where(taskent.CreatorID(params.UserID)).
				WithCreator().
				Order(ent.Asc()).
				First(ctx)
			if err != nil {
				return nil, platform.NotFound("Task is not found")
			}
			task1.Lo = newLoTask.CustomMult
			break
		}
	}

	task1.CustomMult = (task1.Hi + task1.Lo) / 2 * (1 + constant)
	task2.CustomMult *= 1 - constant

	question, err := s.GenerateQuestion(ctx, domain.TaskFromEnt(task1))
	if err != nil {
		return nil, err
	}

	err = s.client.Update().
		Where(taskent.ID(task1.ID)).
		SetCustomMult(task1.CustomMult).
		Exec(ctx)
	if err != nil {
		return nil, platform.WrapInternal(err)
	}

	err = s.client.Update().
		Where(taskent.ID(task2.ID)).
		SetCustomMult(task2.CustomMult).
		Exec(ctx)
	if err != nil {
		return nil, platform.WrapInternal(err)
	}

	return question, nil
	/*type AnswerQuestionDTO struct {
		UserID        int `json:"-"`
		Response      int `json:"response"` // -1,0,1
		TaskID        int `json:"-"`
		CompareTaskID int `json:"task_id"`
	}*/
}

const (
	Epsilon = 0.05
)

func (s Service) GenerateQuestion(ctx context.Context, task *domain.Task) (*domain.Question, error) {
	nearestTask, err := GetNearest(s, ctx, task)
	if err != nil {
		return nil, err
	}
	if nearestTask == nil {
		return nil, nil
	}
	//if task.Hi == 0 || task.Lo == 0 {
	return &domain.Question{
		Text: fmt.Sprintf("Задача %s є важливішою для Вас ніж %s?",
			task.Title, nearestTask.Title),
		CompareTaskID: nearestTask.ID,
	}, nil
	//}

}

func GetNearest(service Service, ctx context.Context, task *domain.Task) (*ent.Task, error) {
	higherTask, err := service.client.Query().
		Where(taskent.CreatorID(task.Creator.ID),
			taskent.CustomMultGTE(task.CustomMult),
			sameRang(task),
		).
		Order(ent.Asc(taskent.FieldCustomMult)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, platform.WrapInternal(err)
	}

	lowerTask, err := service.client.Query().
		Where(taskent.CreatorID(task.Creator.ID),
			taskent.CustomMultLTE(task.CustomMult),
			sameRang(task),
		).
		Order(ent.Desc(taskent.FieldCustomMult)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, platform.WrapInternal(err)
	}

	if higherTask == nil || lowerTask == nil {
		return nil, nil
	}

	if math.Abs(higherTask.CustomMult-task.CustomMult) > Epsilon &&
		math.Abs(lowerTask.CustomMult-task.CustomMult) > Epsilon {
		return nil, nil
	}
	if higherTask.CustomMult-task.CustomMult < task.CustomMult-lowerTask.CustomMult {
		return higherTask, nil
	} else {
		return lowerTask, nil
	}
}

func sameRang(task *domain.Task) predicate.Task {
	return taskent.Or(
		taskent.And(
			taskent.Importance(int8(task.Importance)),
			taskent.Urgency(int8(task.Urgency)),
		),
		taskent.And(
			taskent.Importance(int8(task.Urgency)),
			taskent.Urgency(int8(task.Importance)),
		),
	)
}

func (s Service) Delete(ctx context.Context, taskID int) error {
	n, err := s.client.Delete().
		Where(taskent.ID(taskID)).
		Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return platform.NotFound("Task is not found")
		}
		return platform.WrapInternal(err)
	}

	if n == 0 {
		return platform.NotFound("Task is not found")
	}

	return nil
}
