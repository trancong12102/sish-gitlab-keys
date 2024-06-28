package http_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/trancong12102/sish-gitlab-keys/internal/http"
	httpmocks "github.com/trancong12102/sish-gitlab-keys/mocks/http"
)

var _httpSrvConfig = &http.ServerConfig{
	Addr: ":45566",
}

type HTTPServerTestSuite struct {
	suite.Suite
}

func TestHTTPServer(t *testing.T) { //nolint:paralleltest
	suite.Run(t, new(HTTPServerTestSuite))
}

func (s *HTTPServerTestSuite) TestRun() {
	handler := httpmocks.NewMockServerHandler(s.T())

	srv := http.NewHTTPSrv(handler, _httpSrvConfig)
	go func() {
		time.Sleep(100 * time.Millisecond)

		_ = srv.Shutdown(context.Background())
	}()

	err := srv.Run()
	s.Require().NoError(err)
}

func (s *HTTPServerTestSuite) TestRun_Error() {
	handler := httpmocks.NewMockServerHandler(s.T())

	srv := http.NewHTTPSrv(handler, _httpSrvConfig)
	srvSamePort := http.NewHTTPSrv(handler, _httpSrvConfig)

	go func() {
		err := srv.Run()
		s.NoError(err)

		time.Sleep(100 * time.Millisecond)

		err = srv.Shutdown(context.Background())
		s.NoError(err)

		err = srvSamePort.Shutdown(context.Background())
		s.NoError(err)
	}()

	time.Sleep(100 * time.Millisecond)

	err := srvSamePort.Run()
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "listen and serve http server: ")
}
