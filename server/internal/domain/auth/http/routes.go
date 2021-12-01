package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/domain"
	"server/internal/platform"
)

func RegisterRoutes(r *gin.RouterGroup, service domain.UserService) {
	// todo add routes
	//r.GET("/login", login(service))
}

func JwtAuth(service domain.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Token")
		if tokenStr == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
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
