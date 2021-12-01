package http

import (
	"github.com/gin-gonic/gin"
	"server/internal/domain"
)

func RegisterRoutes(r *gin.RouterGroup, service domain.UserService) {
	//r.GET("/me", meGet(service))

}
