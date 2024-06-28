package http_test

import (
	stdhttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/trancong12102/sish-gitlab-keys/internal/http"
	httpmocks "github.com/trancong12102/sish-gitlab-keys/mocks/http"
	"github.com/trancong12102/sish-gitlab-keys/test/testutil"
)

type RouterTestSuite struct {
	suite.Suite
}

func TestRouter(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(RouterTestSuite))
}

func (s *RouterTestSuite) TestHealthcheckHandler() {
	healthcheckHandler := httpmocks.NewMockRouterHandler(s.T())

	healthcheckHandler.EXPECT().
		ServeHTTP(testutil.AnyThingOfTypeHTTPResponseWriter, testutil.AnyThingOfTypeHTTPRequest).
		Once()

	authHandler := httpmocks.NewMockAuthHandler(s.T())

	router := http.NewRouter(authHandler, healthcheckHandler)

	req := httptest.NewRequest(stdhttp.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	s.Require().Equal(stdhttp.StatusOK, w.Code)
}

func (s *RouterTestSuite) TestAuthHandler() {
	authHandler := httpmocks.NewMockAuthHandler(s.T())

	authHandler.EXPECT().
		AuthorizePubKey(testutil.AnyThingOfTypeHTTPResponseWriter, testutil.AnyThingOfTypeHTTPRequest).
		Once()

	healthcheckHandler := httpmocks.NewMockRouterHandler(s.T())

	router := http.NewRouter(authHandler, healthcheckHandler)

	req := httptest.NewRequest(stdhttp.MethodPost, "/auth", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	s.Require().Equal(stdhttp.StatusOK, w.Code)
}
