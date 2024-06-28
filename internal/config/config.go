package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Environment string

const (
	EnvironmentTypeDevelopment Environment = "development"
	EnvironmentTypeProduction  Environment = "production"
)

//go:generate go run github.com/g4s8/envdoc@latest -output ../../docs/environments.md -type Config
type Config struct {
	// Specify the environment: development or production
	AppEnv Environment `env:"APP_ENV" envDefault:"development"`
	// Address to listen on
	ListenAddr string `env:"LISTEN_ADDR" envDefault:":8080"`
	// Gitlab URL
	GitlabURL string `env:"GITLAB_URL,notEmpty"`
	// Gitlab Access Token
	GitlabAccessToken string `env:"GITLAB_ACCESS_TOKEN,notEmpty"`
}

func LoadConfig() (*Config, error) {
	var config Config

	err := env.ParseWithOptions(&config, env.Options{
		RequiredIfNoDef:       true,
		UseFieldNameByDefault: true,
	})
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	return &config, nil
}
