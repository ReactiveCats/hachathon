package bootstrap

import "server/internal/config"

type Dependencies struct {
	Config config.Config
	//EntClient *ent.Client
	Server Server
}

func NewDependencies(config config.Config) Dependencies {
	return Dependencies{Config: config}
}
