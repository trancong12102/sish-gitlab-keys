package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RouterTestSuite struct {
	suite.Suite
}

func TestRouter(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(RouterTestSuite))
}

func (s *RouterTestSuite) TestHealthcheckHandler() {
	healthcheckHandler := NewMockHTTPHandler(s.T())

	healthcheckHandler.EXPECT().
		ServeHTTP(anyThingOfTypeHTTPResponseWriter, anyThingOfTypeHTTPRequest).
		Once()

	authHandler := NewMockAuthHandler(s.T())

	router := NewRouter(authHandler, healthcheckHandler)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	s.Require().Equal(http.StatusOK, w.Code)
}

func (s *RouterTestSuite) TestAuthHandler() {
	authHandler := NewMockAuthHandler(s.T())

	authHandler.EXPECT().
		AuthorizePubKey(anyThingOfTypeHTTPResponseWriter, anyThingOfTypeHTTPRequest).
		Once()

	healthcheckHandler := NewMockHTTPHandler(s.T())

	router := NewRouter(authHandler, healthcheckHandler)

	req := httptest.NewRequest(http.MethodPost, "/auth", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	s.Require().Equal(http.StatusOK, w.Code)
}
