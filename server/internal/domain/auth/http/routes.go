package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/config"
	"server/internal/domain"
	"server/internal/platform"
)

func RegisterRoutes(r *gin.RouterGroup, service domain.UserService) {
	r.GET("/login", login(service))
	r.GET("/signup", signup(service))
}

func JwtAuth(config config.Config, service domain.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Token")
		if tokenStr == "" {
			if !config.DefaultAuthUser {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			tokenStr, _ = service.JWTToken(&domain.User{ID: 1})
		}

		userID, err := service.DataFromJWT(tokenStr)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user *domain.User

		user, err = service.ByID(ctx.Request.Context(), domain.GetUserDTO{ID: userID})
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set(platform.UserCtxKey, user)

		ctx.Next()
	}
}

// login godoc
// @Summary 	Login
// @Description Login
// @Tags 		auth
// @ID 			login
// @Produce  	json
// @Param   	username  	query  string  	true  "username"
// @Success 	200 {object} string
// @Router 		/auth/login [get]
func login(service domain.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Query("username")

		token, err := service.Login(ctx.Request.Context(), username)
		if err != nil {
			platform.GinErrResponse(ctx, err)
			return
		}

		platform.GinOkResponse(ctx, http.StatusOK, token)
	}
}

// signup godoc
// @Summary 	Signup
// @Description Signup
// @Tags 		auth
// @ID 			signup
// @Produce  	json
// @Param   	username  	query  string  	true  "username"
// @Success 	200 {object} string
// @Router 		/auth/signup [get]
func signup(service domain.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Query("username")

		token, err := service.Signup(ctx.Request.Context(), username)
		if err != nil {
			platform.GinErrResponse(ctx, err)
			return
		}

		platform.GinOkResponse(ctx, http.StatusOK, token)
	}
}
