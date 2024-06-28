//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/hellofresh/health-go/v5"

	"github.com/trancong12102/sish-gitlab-keys/internal/auth"
	"github.com/trancong12102/sish-gitlab-keys/internal/gitlab"
	"github.com/trancong12102/sish-gitlab-keys/internal/http"
)

func InitializeHTTPSrv(
	httpSrvConfig *http.ServerConfig,
	gitlabClientConfig *gitlab.ClientConfig,
	healthcheckOpts ...health.Option,
) (*http.Server, error) {
	wire.Build(
		http.NewHTTPSrv,
		http.NewRouter,
		auth.NewHandler,
		http.NewHealthcheck,
		auth.NewService,
		gitlab.NewClient,
		wire.Bind(new(http.ServerHandler), new(*http.Router)),
		wire.Bind(new(http.RouterHandler), new(*http.Healthcheck)),
		wire.Bind(new(http.AuthHandler), new(*auth.Handler)),
		wire.Bind(new(auth.PubKeyGetter), new(*gitlab.Client)),
		wire.Bind(new(auth.Authorizer), new(*auth.Service)),
	)

	return &http.Server{}, nil
}
