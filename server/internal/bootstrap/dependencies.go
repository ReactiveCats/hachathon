package bootstrap

import (
	"context"
	"server/internal/config"
	"server/internal/domain"
	"server/internal/ent"
)

type Dependencies struct {
	Config    config.Config
	EntClient *ent.Client

	UserService domain.UserService
	TaskService domain.TaskService
}

func NewDependencies(config config.Config, entClient *ent.Client,
	user domain.UserService,
	task domain.TaskService) Dependencies {
	return Dependencies{
		Config:      config,
		EntClient:   entClient,
		UserService: user,
		TaskService: task,
	}
}

// NewEntClient wraps the creation of ent.Client
func NewEntClient(cfg config.Config) (*ent.Client, error) {
	client, clientErr := ent.Open(cfg.Database.Driver, cfg.Database.URL)
	if clientErr != nil {
		return nil, clientErr
	}

	return client, client.Schema.Create(context.Background())
}
