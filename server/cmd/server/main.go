// @title HackathonAPI
// @version 0.1
// @description API for hackathon project
//
// @host localhost:8080
// @BasePath /api
// @schemes http
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @tag.name 			auth
// @tag.description  	Authorization endpoints
// @tag.name 			tasks
// @tag.description  	Tasks endpoints
// @tag.name 			tags
// @tag.description  	Tags endpoints

package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"server/internal/bootstrap"
)

func main() {
	logrus.Infof("Starting api")

	deps, err := bootstrap.Up(context.Background())
	if err != nil {
		panic(err)
	}

	server := bootstrap.NewServer(deps)
	err = server.Start()
	if err != nil {
		panic(err)
	}
}
