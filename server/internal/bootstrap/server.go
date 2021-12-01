package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	authHTTP "server/internal/domain/auth/http"
	userHTTP "server/internal/domain/user/http"
)

type Server struct {
	server *http.Server
}

func (s Server) Start() error {
	return s.server.ListenAndServe()
}

func NewServer(deps Dependencies) *Server {
	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger())

	// мб это нафиг не нужно, но ща час ночи и мне лень проверять,
	// уберите, если джин ошибки и без того принтит
	engine.Use(func(ctx *gin.Context) {
		ctx.Next()
		for i, s := range ctx.Errors.Errors() {
			logrus.Errorf("Error #%02d: %s", i+1, s)
		}
	})

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	handler := engine.Group("/api")

	authHTTP.RegisterRoutes(handler.Group("/auth"), deps.UserService)

	handlerAuthorized := handler.Group("")
	handlerAuthorized.Use(authHTTP.JwtAuth(deps.UserService))

	userHTTP.RegisterRoutes(handlerAuthorized.Group("/users"), deps.UserService)

	return &Server{
		server: &http.Server{
			Addr:    deps.Config.Server.Addr,
			Handler: engine,
		},
	}
}
