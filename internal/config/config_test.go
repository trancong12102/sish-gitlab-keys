package config_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/trancong12102/sish-gitlab-keys/internal/config"
)

type ConfigTestSuite struct {
	suite.Suite
}

func TestConfig(t *testing.T) { //nolint:paralleltest
	suite.Run(t, new(ConfigTestSuite))
}

func (s *ConfigTestSuite) TestLoadConfig() {
	s.T().Setenv("APP_ENV", "production")
	s.T().Setenv("LISTEN_ADDR", ":8081")
	s.T().Setenv("GITLAB_URL", "https://example.com")
	s.T().Setenv("GITLAB_ACCESS_TOKEN", "token")

	cfg, err := config.LoadConfig()

	s.Require().NoError(err)
	s.Require().Equal(&config.Config{
		AppEnv:            config.EnvironmentTypeProduction,
		ListenAddr:        ":8081",
		GitlabURL:         "https://example.com",
		GitlabAccessToken: "token",
	}, cfg)
}

func (s *ConfigTestSuite) TestLoadConfig_InvalidEnv() {
	s.T().Setenv("APP_ENV", "invalid")
	s.T().Setenv("LISTEN_ADDR", ":8081")
	s.T().Setenv("GITLAB_URL", "")
	s.T().Setenv("GITLAB_ACCESS_TOKEN", "")

	cfg, err := config.LoadConfig()

	s.Require().Error(err)
	s.Require().Nil(cfg)
}
