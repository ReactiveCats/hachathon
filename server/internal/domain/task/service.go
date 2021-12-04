package task

import (
	"context"
	"server/internal/config"
	"server/internal/domain"
	"server/internal/ent"
	taskent "server/internal/ent/task"
	userent "server/internal/ent/user"
	"server/internal/platform"
	"time"
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
	if dto.Complexity != nil {
		query.Where(taskent.ComplexityEQ(int8(*dto.Complexity)))
	}
	if dto.Priority != nil {
		query.Where(taskent.PriorityEQ(int8(*dto.Priority)))
	}
	if dto.OrderBy != nil && dto.Order != nil {
		if *dto.Order == "asc" {
			query.Order(ent.Asc(*dto.OrderBy))
		} else {
			query.Order(ent.Desc(*dto.OrderBy))
		}
	}
	query.Where(taskent.HasCreatorWith(userent.ID(dto.UserID)))
	tasks, err := query.All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, platform.NotFound("Tasks are not found")
		}
		return nil, platform.WrapInternal(err)
	}
	return domain.TasksFromEnt(tasks), nil
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
		SetComplexity(int8(taskDTO.Complexity)).
		SetPriority(int8(taskDTO.Complexity)).
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
		SetPriority(int8(taskDTO.Priority)).
		SetComplexity(int8(taskDTO.Complexity)).
		SetNillableDeadline(taskDTO.Deadline).
		SetNillableEstimated(taskDTO.Estimated).
		SetCreatorID(taskDTO.UserID).
		SetF(0). // костыли, которые мы заслужили
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, nil, platform.NotFound("User not found")
		}
		return nil, nil, platform.WrapInternal(err)
	}

	task := domain.TaskFromEnt(entTask)

	f := CalculateF(task)
	task.F = f

	err = s.client.Update().
		Where(taskent.ID(task.ID)).
		SetF(f).
		Exec(ctx)
	if err != nil {
		return nil, nil, platform.WrapInternal(err)
	}

	question, err := s.AskQuestion(ctx, task)
	if err != nil {
		return nil, nil, err
	}

	return task, question, nil
}

func CalculateF(task *domain.Task) float64 {
	timeLeft := task.Deadline.Unix() - time.Now().Unix()
	f := float64(task.Estimated) / float64(timeLeft)

	return f
}

func (s Service) AnswerQuestion(ctx context.Context, params domain.AnswerQuestionDTO) (*domain.Question, error) {
	panic("implement me")
}

func (s Service) AskQuestion(ctx context.Context, task *domain.Task) (*domain.Question, error) {
	panic("implement me")
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
