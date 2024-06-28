package http

import (
	"fmt"
	"net/http"

	"github.com/hellofresh/health-go/v5"
)

const (
	_serviceName    = "sish-gitlab-keys"
	_serviceVersion = "1"
)

var _ http.Handler = (*Healthcheck)(nil)

type Healthcheck struct {
	*health.Health
}

func (h *Healthcheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Health.Handler().ServeHTTP(w, r)
}

func NewHealthcheck(opts ...health.Option) (*Healthcheck, error) {
	opts = append(opts, health.WithComponent(health.Component{
		Name:    _serviceName,
		Version: _serviceVersion,
	}))

	healthcheck, err := health.New(opts...)
	if err != nil {
		return nil, fmt.Errorf("create healthcheck: %w", err)
	}

	return &Healthcheck{
		Health: healthcheck,
	}, nil
}
