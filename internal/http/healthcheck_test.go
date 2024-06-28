package http_test

import (
	stdhttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/hellofresh/health-go/v5"
	"github.com/stretchr/testify/suite"

	"github.com/trancong12102/sish-gitlab-keys/internal/http"
)

type HealthcheckTestSuite struct {
	suite.Suite
}

func TestHealthcheck(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(HealthcheckTestSuite))
}

func (s *HealthcheckTestSuite) TestNewHealthcheckHandler() {
	healthcheck, err := http.NewHealthcheck()
	s.Require().NoError(err)
	s.Require().NotNil(healthcheck)
}

func (s *HealthcheckTestSuite) TestHealthcheckHandler_Error() {
	healthcheck, err := http.NewHealthcheck(
		health.WithChecks(health.Config{}),
	)
	s.Require().Nil(healthcheck)
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "create healthcheck: ")
}

func (s *HealthcheckTestSuite) TestHealthcheckHandler() {
	healthcheck, err := http.NewHealthcheck()
	s.Require().NoError(err)
	s.Require().NotNil(healthcheck)

	req := httptest.NewRequest(stdhttp.MethodPost, "/auth", nil)
	w := httptest.NewRecorder()

	healthcheck.ServeHTTP(w, req)
	s.Require().Equal(stdhttp.StatusOK, w.Code)
}
