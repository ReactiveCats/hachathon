//go:build wireinject
// +build wireinject

package bootstrap

import (
	"context"
	"github.com/google/wire"
	"server/internal/config"
	"server/internal/domain"
	"server/internal/domain/task"
	"server/internal/domain/user"
)

func Up(ctx context.Context) (Dependencies, error) {
	wire.Build(
		NewDependencies,
		config.New,
		NewEntClient,
		user.NewService,
		wire.Bind(new(domain.UserService), new(*user.Service)),
		task.NewService,
		wire.Bind(new(domain.TaskService), new(*task.Service)),
	)
	return Dependencies{}, nil
}
