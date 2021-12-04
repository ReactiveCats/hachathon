package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/config"
	"server/internal/domain"
	"server/internal/platform"
	"strings"
)

func RegisterRoutes(r *gin.RouterGroup, service domain.UserService) {
	r.POST("/login", login(service))
	r.POST("/signup", signup(service))
}

func JwtAuth(config config.Config, service domain.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authStr := ctx.GetHeader("Authorization")
		if authStr == "" {
			if !config.DefaultAuthUser {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			authStr, _ = service.JWTToken(&domain.User{ID: 1})
			authStr = "Bearer " + authStr
		}

		token := strings.Replace(authStr, "Bearer ", "", 1)

		userID, err := service.DataFromJWT(token)
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
// @Param   	username  	body  domain.LoginDTO  	true  "username"
// @Success 	200 {object} string
// @Router 		/auth/login [post]
func login(service domain.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params domain.LoginDTO

		err := ctx.ShouldBind(&params)
		if err != nil {
			platform.GinErrResponse(ctx, platform.WrapUnprocessableEntity(err, "username is empty"))
			return
		}

		token, err := service.Login(ctx.Request.Context(), params.Username)
		if err != nil {
			platform.GinErrResponse(ctx, err)
			return
		}

		platform.GinOkResponse(ctx, http.StatusOK, domain.AccessToken{AccessToken: token})
	}
}

// signup godoc
// @Summary 	Signup
// @Description Signup
// @Tags 		auth
// @ID 			signup
// @Produce  	json
// @Param   	username  	body  domain.SignupDTO  	true  "username"
// @Success 	200 {object} string
// @Router 		/auth/signup [post]
func signup(service domain.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params domain.SignupDTO

		err := ctx.ShouldBind(&params)
		if err != nil {
			platform.GinErrResponse(ctx, platform.WrapUnprocessableEntity(err, "username is empty"))
			return
		}

		token, err := service.Signup(ctx.Request.Context(), params.Username)
		if err != nil {
			platform.GinErrResponse(ctx, err)
			return
		}

		platform.GinOkResponse(ctx, http.StatusOK, domain.AccessToken{AccessToken: token})
	}
}
