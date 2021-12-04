package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/domain"
	"server/internal/platform"
	"strconv"
)

func RegisterRoutes(r *gin.RouterGroup, service domain.TagService) {
	r.GET("/", getTags(service))
	r.GET("/:tag_id", getTagByID(service))
	r.PUT("/:tag_id", putTag(service))
	r.POST("", createTag(service))
	r.DELETE("/:tag_id", deleteTag(service))
}

// getTags godoc
// @Summary 	Get all user's tags
// @Description Get all user's tags
//// @Security 	ApiKeyAuth
// @Tags 		tags
// @ID 			get_tags
// @Produce  	json
// @Success 	200 {array} domain.Tag
// @Router 		/tag [get]
func getTags(service domain.TagService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user := ctx.Value(platform.UserCtxKey).(*domain.User)

		tags, err := service.All(ctx.Request.Context(), domain.AllTagsDTO{UserID: user.ID})
		if err != nil {
			platform.GinErrResponse(ctx, err)
			return
		}

		platform.GinOkResponse(ctx, http.StatusOK, tags)
	}
}

// getTagByID godoc
// @Summary 	Get tag by id
// @Description Get tag by id
//// @Security 	ApiKeyAuth
// @Tags 		tags
// @ID 			get_tag_by_id
// @Param 		tag_id 	path 	int		true 	"tag id"
// @Produce  	json
// @Success 	200 {object} domain.Tag
// @Router 		/tag/{tag_id} [get]
func getTagByID(service domain.TagService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user := ctx.Value(platform.UserCtxKey).(*domain.User)

		tagID, err := strconv.Atoi(ctx.Param("tag_id"))
		if err != nil {
			platform.GinErrResponse(ctx, platform.NotFound("Tag is not found"))
			return
		}

		tag, err := service.ByID(ctx.Request.Context(), domain.TagsByIDDTO{
			UserID: user.ID,
			TagID:  tagID,
		})
		if err != nil {
			platform.GinErrResponse(ctx, err)
			return
		}

		platform.GinOkResponse(ctx, http.StatusOK, tag)
	}

}

// putTag godoc
// @Summary 	Edit tag
// @Description Edit tag
//// @Security 	ApiKeyAuth
// @Tags 		tags
// @ID 			put_tag
// @Param 		tag_id 		path 	string				true 	"tag id"
// @Param 		tag 		body 	domain.TagPutDTO	true 	"tag object"
// @Produce  	json
// @Success 	200 {object} domain.Tag
// @Router 		/tag/{tag_id} [put]
func putTag(service domain.TagService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user := ctx.Value(platform.UserCtxKey).(*domain.User)

		tagID, err := strconv.Atoi(ctx.Param("tag_id"))
		if err != nil {
			platform.GinErrResponse(ctx, platform.NotFound("Tag is not found"))
			return
		}

		var params = domain.TagPutDTO{
			UserID: user.ID,
			TagID:  tagID,
		}

		err = ctx.ShouldBind(&params)
		if err != nil {
			platform.GinErrResponse(ctx, platform.WrapUnprocessableEntity(err, "wrong tag"))
			return
		}

		tag, err := service.Update(ctx.Request.Context(), params)
		if err != nil {
			platform.GinErrResponse(ctx, err)
			return
		}

		platform.GinOkResponse(ctx, http.StatusOK, tag)
	}
}

// postTag godoc
// @Summary 	Add new tag
// @Description Add new tag
//// @Security 	ApiKeyAuth
// @Tags 		tags
// @ID 			post_tag
// @Produce  	json
// @Success 	200 {object} domain.Tag
// @Router 		/tag [post]
func createTag(service domain.TagService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user := ctx.Value(platform.UserCtxKey).(*domain.User)

		var tagDTO = domain.CreateTagDTO{
			UserID: user.ID,
		}
		err := ctx.ShouldBind(&tagDTO)
		if err != nil {
			platform.GinErrResponse(ctx, platform.UnprocessableEntity("wrong tag entity"))
			return
		}

		tag, err := service.Create(ctx.Request.Context(), tagDTO)
		if err != nil {
			platform.GinErrResponse(ctx, err)
			return
		}

		platform.GinOkResponse(ctx, http.StatusCreated, tag)
	}
}

// deleteTag godoc
// @Summary 	Delete tag
// @Description Delete tag
//// @Security 	ApiKeyAuth
// @Tags 		tags
// @ID 			delete_tag
// @Param 		tag_id 		path 	string		true 	"tag id"
// @Produce  	json
// @Success 	204
// @Router 		/tag/{tag_id} [delete]
func deleteTag(service domain.TagService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		user := ctx.Value(platform.UserCtxKey).(*domain.User)

		tagID, err := strconv.Atoi(ctx.Param("tag_id"))
		if err != nil {
			platform.GinErrResponse(ctx, platform.NotFound("Tag is not found"))
			return
		}

		err = service.Delete(ctx.Request.Context(), domain.DeleteTagDTO{
			UserID: user.ID,
			TagID:  tagID,
		})
		if err != nil {
			platform.GinErrResponse(ctx, err)
			return
		}

		platform.GinOkResponse(ctx, http.StatusNoContent)
	}
}
