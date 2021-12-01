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
}

func NewDependencies(config config.Config, entClient *ent.Client, user domain.UserService) Dependencies {
	return Dependencies{
		Config:      config,
		EntClient:   entClient,
		UserService: user,
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
