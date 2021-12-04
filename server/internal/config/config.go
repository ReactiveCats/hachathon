package config

import (
	"github.com/spf13/viper"
	"io/fs"
	"os"
)

type Config struct {
	Database        Database
	Server          Server
	Jwt             Jwt
	DefaultAuthUser bool
}

type Database struct {
	Driver string
	URL    string
}

type Server struct {
	LogLevel string
	Addr     string
}

type Jwt struct {
	Secret string
}

const _configDefaultFilename = "api_config.yml"

func New() (Config, error) {
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	cfgFile := os.Getenv("API_CONFIG")
	if cfgFile == "" {
		cfgFile = _configDefaultFilename
	}

	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(*fs.PathError); !ok {
			return Config{}, err
		}
	}

	return Config{
		Database: Database{
			Driver: viper.GetString("data_driver"),
			URL:    viper.GetString("data_url"),
		},
		Server: Server{
			LogLevel: viper.GetString("log_level"),
			Addr:     viper.GetString("bind_addr"),
		},
		Jwt: Jwt{
			Secret: viper.GetString("jwt_secret"),
		},
		DefaultAuthUser: viper.GetBool("default_auth_user"),
	}, nil
}
