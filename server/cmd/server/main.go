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
// @name Token

// @tag.name 			auth
// @tag.description  	Authorization endpoints

package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"math/rand"
	"server/internal/bootstrap"
	"time"
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

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
