package http

import (
	"github.com/gin-gonic/gin"
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
func getTasks(domain.TaskService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

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
// @Param 		task_id 	path 	string		true 	"task id"
// @Produce  	json
// @Success 	200 {object} domain.Task
// @Router 		/task/{task_id} [put]
func putTask(domain.TaskService) func(ctx *gin.Context) {
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
