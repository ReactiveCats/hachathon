package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"server/internal/domain"
	"server/internal/platform"
	"strconv"
)

func RegisterRoutes(r *gin.RouterGroup, service domain.TaskService) {
	r.GET("/", getTasks(service))
	r.GET("/:task_id", getTaskByID(service))
	r.PUT("/:task_id", putTask(service))
	r.POST("", postTask(service))
	r.DELETE("/:task_id", deleteTask(service))
}

// getTasks godoc
// @Summary 	Get all user tasks
// @Description Get all user tasks
//// @Security 	ApiKeyAuth
// @Tags 		tasks
// @ID 			get_tasks
// @Produce  	json
// @Success 	200 {array} domain.Task
// @Router 		/task [get]
func getTasks(service domain.TaskService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user := ctx.Value(platform.UserCtxKey).(*domain.User)
		dto := domain.GetTaskDTO{
			UserID: user.ID,
		}

		estimated, ok := ctx.Get("estimated")
		if ok {
			parsedEstimated, err := strconv.Atoi(estimated.(string))
			if err != nil {
				platform.GinErrResponse(ctx, platform.Conflict("Invalid data"))
				return
			}
			dto.Estimated = &parsedEstimated
		}

		complexity, ok := ctx.Get("complexity")
		if ok {
			parsedComplexity := complexity.(string)
			dto.Complexity = &parsedComplexity
		}

		priority, ok := ctx.Get("priority")
		if ok {
			parsedPriority := priority.(string)
			dto.Priority = &parsedPriority
		}

		order, ok := ctx.Get("order")
		if ok {
			parsedOrder := order.(string)
			dto.Order = &parsedOrder
		}

		orderBy, ok := ctx.Get("order_by")
		if ok {
			parsedOrderBy := orderBy.(string)
			dto.OrderBy = &parsedOrderBy
		}

		err := binding.Validator.ValidateStruct(&dto)
		if err != nil {
			platform.GinErrResponse(ctx, platform.Conflict("Invalid data"))
			return
		}

		tasks, err := service.Fetch(ctx, dto)
		if err != nil {
			return
		}

		ctx.JSON(http.StatusOK, tasks)
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
	}
}

// postTasks godoc
// @Summary 	Add new task
// @Description Add new task
//// @Security 	ApiKeyAuth
// @Tags 		tasks
// @ID 			post_task
// @Produce  	json
// @Success 	200 {object} domain.Task
// @Router 		/task [post]
func postTask(service domain.TaskService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
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
