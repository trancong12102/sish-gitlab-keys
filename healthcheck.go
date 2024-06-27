package main

import (
	"fmt"
	"net/http"

	"github.com/hellofresh/health-go/v5"
)

const (
	_serviceName    = "sish-gitlab-keys"
	_serviceVersion = "1"
)

func NewHealthCheckHandler() (http.Handler, error) {
	healthcheck, err := health.New(health.WithComponent(health.Component{
		Name:    _serviceName,
		Version: _serviceVersion,
	}))
	if err != nil {
		return nil, fmt.Errorf("create healthcheck: %w", err)
	}

	return healthcheck.Handler(), nil
}
