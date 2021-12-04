package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"server/internal/domain"
	"server/internal/platform"
	"strconv"
	"strings"
	"time"
)

func RegisterRoutes(r *gin.RouterGroup, service domain.TaskService) {
	r.GET("/", getTasks(service))
	r.GET("/:task_id", getTaskByID(service))
	r.PUT("/:task_id", putTask(service))
	r.POST("", postTask(service))
	r.POST("/:task_id/question", answerQuestion(service))
	r.DELETE("/:task_id", deleteTask(service))
}

// getTasks godoc
// @Summary 	Get all user tasks
// @Description Get all user tasks
//// @Security 	ApiKeyAuth
// @Tags 		tasks
// @ID 			get_tasks
// @Param 		estimated 	query 	string		false 	"estimated"
// @Param 		importance 	query 	string		false 	"importance"
// @Param 		urgency 	query 	string		false 	"urgency"
// @Param 		order 		query 	string		false 	"order field (asc/desc)"
// @Param 		orderBy 	query 	string		false 	"order field (e.g. deadline, estimated, importance)"
// @Produce  	json
// @Success 	200 {array} domain.Task
// @Router 		/task [get]
func getTasks(service domain.TaskService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user := ctx.Value(platform.UserCtxKey).(*domain.User)
		dto := domain.GetTaskDTO{
			UserID: user.ID,
		}

		estimated, ok := ctx.GetQuery("estimated")
		if ok {
			parsedEstimated, err := strconv.Atoi(estimated)
			if err != nil {
				platform.GinErrResponse(ctx, platform.Conflict("invalid estimated (must be int)"))
				return
			}
			dto.Estimated = &parsedEstimated
		}

		importance, ok := ctx.GetQuery("importance")
		if ok {
			parsedImportance, err := strconv.Atoi(importance)
			if err != nil {
				platform.GinErrResponse(ctx, platform.Conflict("invalid importance (must be int)"))
				return
			}
			dto.Importance = &parsedImportance
		}

		urgency, ok := ctx.GetQuery("urgency")
		if ok {
			parsedUrgency, err := strconv.Atoi(urgency)
			if err != nil {
				platform.GinErrResponse(ctx, platform.Conflict("invalid urgency (must be int)"))
				return
			}
			dto.Urgency = &parsedUrgency
		}

		order, ok := ctx.GetQuery("order")
		if ok {
			dto.Order = &order
		}

		orderBy, ok := ctx.GetQuery("orderBy")
		if ok {
			dto.OrderBy = &orderBy
		}

		err := binding.Validator.ValidateStruct(&dto)
		if err != nil {
			platform.GinErrResponse(ctx, platform.UnprocessableEntity("Invalid data"))
			return
		}

		tasks, err := service.Fetch(ctx.Request.Context(), dto)
		if err != nil {
			platform.GinErrResponse(ctx, err)
			return
		}

		platform.GinOkResponse(ctx, http.StatusOK, tasks)
	}
}

// getTaskByID godoc
// @Summary 	Get task by id
// @Description Get task by id
//// @Security 	ApiKeyAuth
// @Tags 		tasks
// @ID 			get_task_by_id
// @Param 		task_id 	path 	int		true 	"task id"
// @Produce  	json
// @Success 	200 {object} domain.Task
// @Router 		/task/{task_id} [get]
func getTaskByID(service domain.TaskService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		taskID, err := strconv.Atoi(ctx.Param("task_id"))
		if err != nil {
			platform.GinErrResponse(ctx, platform.NotFound("Task is not found"))
			return
		}

		task, err := service.ByID(ctx.Request.Context(), taskID)
		if err != nil {
			platform.GinErrResponse(ctx, err)
			return
		}

		platform.GinOkResponse(ctx, http.StatusOK, task)
	}

}

// answerQuestion godoc
// @Summary 	Answer question
// @Description Answer question to make F more precise
//// @Security 	ApiKeyAuth
// @Tags 		tasks
// @ID 			answer_task_question
// @Param 		task_id 	path 	string						true 	"task id"
// @Param 		answer 		body 	domain.AnswerQuestionDTO	true 	"answer question object"
// @Produce  	json
// @Success 	200 {object} domain.Question
// @Router 		/task/{task_id}/question [post]
func answerQuestion(service domain.TaskService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user := ctx.Value(platform.UserCtxKey).(*domain.User)

		taskID, err := strconv.Atoi(ctx.Param("task_id"))
		if err != nil {
			platform.GinErrResponse(ctx, platform.NotFound("Task is not found"))
			return
		}

		var params = domain.AnswerQuestionDTO{
			UserID:        user.ID,
			TaskID:        taskID,
			CompareTaskID: taskID,
		}

		err = ctx.ShouldBind(&params)
		if err != nil {
			platform.GinErrResponse(ctx, platform.WrapUnprocessableEntity(err, "wrong task entity"))
			return
		}

		question, err := service.AnswerQuestion(ctx.Request.Context(), params)
		if err != nil {
			platform.GinErrResponse(ctx, err)
			return
		}

		platform.GinOkResponse(ctx, http.StatusOK, question)
	}
}

// putTask godoc
// @Summary 	Edit task
// @Description Edit task
//// @Security 	ApiKeyAuth
// @Tags 		tasks
// @ID 			put_task
// @Param 		task_id 	path 	string				true 	"task id"
// @Param 		task 		body 	domain.TaskPutDTO	true 	"task object"
// @Produce  	json
// @Success 	200 {object} domain.Task
// @Router 		/task/{task_id} [put]
func putTask(service domain.TaskService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user := ctx.Value(platform.UserCtxKey).(*domain.User)

		taskID, err := strconv.Atoi(ctx.Param("task_id"))
		if err != nil {
			platform.GinErrResponse(ctx, platform.NotFound("Task is not found"))
			return
		}

		var params domain.TaskPutDTO

		err = ctx.ShouldBind(&params)
		if err != nil {
			platform.GinErrResponse(ctx, platform.WrapUnprocessableEntity(err, "wrong task"))
			return
		}

		if params.DeadlineDateStr != nil {
			deadline, err := time.Parse(time.RFC3339, *params.DeadlineDateStr)
			if err != nil {
				platform.GinErrResponse(ctx, platform.WrapUnprocessableEntity(err, "wrong date"))
				return
			}
			params.Deadline = &deadline
		}

		params.TaskID = taskID
		params.UserID = user.ID

		task, err := service.Update(ctx.Request.Context(), params)
		if err != nil {
			platform.GinErrResponse(ctx, err)
			return
		}

		platform.GinOkResponse(ctx, http.StatusOK, task)
	}
}

// postTasks godoc
// @Summary 	Add new task
// @Description Add new task
//// @Security 	ApiKeyAuth
// @Tags 		tasks
// @ID 			post_task
// @Param 		task 	body 	domain.CreateTaskDTO	true 	"task object"
// @Produce  	json
// @Success 	200 {object} domain.CreateTaskAnswer
// @Router 		/task [post]
func postTask(service domain.TaskService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user := ctx.Value(platform.UserCtxKey).(*domain.User)

		var taskDTO domain.CreateTaskDTO
		err := ctx.ShouldBind(&taskDTO)
		if err != nil {
			switch err.(type) {
			case validator.ValidationErrors:
				var sb strings.Builder
				sb.WriteString("Error: Field validation failed for: ")
				for _, fieldErr := range err.(validator.ValidationErrors) {
					sb.WriteString(fmt.Sprintf("%s, ", fieldErr.Field()))
				}
				platform.GinErrResponse(ctx, platform.UnprocessableEntity(sb.String()))
				return
			case *time.ParseError:
				platform.GinErrResponse(ctx, platform.UnprocessableEntity("Incorrect date format"))
				return
			default:
				platform.GinErrResponse(ctx, platform.UnprocessableEntity(err.Error()))
				return
			}

		}

		taskDTO.UserID = user.ID

		task, question, err := service.Create(ctx.Request.Context(), taskDTO)
		if err != nil {
			platform.GinErrResponse(ctx, err)
			return
		}

		platform.GinOkResponse(ctx, http.StatusCreated, domain.CreateTaskAnswer{Task: task, Question: question})
	}
}

// deleteTask godoc
// @Summary 	Get list of tasks
// @Description Get list of tasks
//// @Security 	ApiKeyAuth
// @Tags 		tasks
// @ID 			delete_task
// @Param 		task_id 	path 	string		true 	"task id"
// @Produce  	json
// @Success 	204
// @Router 		/task/{task_id} [delete]
func deleteTask(service domain.TaskService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		taskID, err := strconv.Atoi(ctx.Param("task_id"))
		if err != nil {
			platform.GinErrResponse(ctx, platform.NotFound("Task is not found"))
			return
		}

		err = service.Delete(ctx.Request.Context(), taskID)
		if err != nil {
			platform.GinErrResponse(ctx, err)
			return
		}

		platform.GinOkResponse(ctx, http.StatusNoContent)
	}
}
