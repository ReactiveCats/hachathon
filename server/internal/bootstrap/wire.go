//go:build wireinject
// +build wireinject

package bootstrap

import (
	"context"
	"github.com/google/wire"
	"server/internal/config"
)

func Up(ctx context.Context) (Dependencies, error) {
	wire.Build(
		NewDependencies,
		config.New,
		//NewEntClient,
		//order.New,
		//wire.Bind(new(domain.OrderService), new(*order.Service)),
	)
	return Dependencies{}, nil
}
