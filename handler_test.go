package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
}

func TestHandler(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(HandlerTestSuite))
}

const (
	reqURL = "https://example.com"
)

var (
	validBody        = []byte(`{"auth_key":"` + validPubKey + `"}`)
	notFoundUserKey  = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIP1ypONGj/C1D/MmsLFNoDUAz7pbqOOkKfJmIqoZMKIv nopain@Nopain-MacBook-Pro.local" //nolint:lll
	notFoundUserBody = []byte(`{"auth_key":"` + notFoundUserKey + `"}`)
	invalidBody      = []byte("xxx")
)

func (s *HandlerTestSuite) TestAuthorizePubKey() {
	service := NewMockAuthorizer(s.T())

	service.EXPECT().
		AuthorizePubKey(anythingOfTypeContext, validPubKey).
		Once().
		Return(nil)

	handler := NewHandler(service)

	body := bytes.NewReader(validBody)
	req := httptest.NewRequest(http.MethodGet, reqURL, body)

	w := httptest.NewRecorder()

	handler.AuthorizePubKey(w, req)
	s.Require().Equal(http.StatusOK, w.Code)
}

func (s *HandlerTestSuite) TestAuthorizePubKey_NotFoundUser() {
	service := NewMockAuthorizer(s.T())

	service.EXPECT().
		AuthorizePubKey(anythingOfTypeContext, notFoundUserKey).
		Once().
		Return(ErrUserNotFound)

	handler := NewHandler(service)

	body := bytes.NewReader(notFoundUserBody)
	req := httptest.NewRequest(http.MethodGet, reqURL, body)

	w := httptest.NewRecorder()

	handler.AuthorizePubKey(w, req)
	s.Require().Equal(http.StatusUnauthorized, w.Code)
}

func (s *HandlerTestSuite) TestAuthorizePubKey_InvalidBody() {
	service := NewMockAuthorizer(s.T())

	handler := NewHandler(service)

	body := bytes.NewReader(invalidBody)
	req := httptest.NewRequest(http.MethodGet, reqURL, body)

	w := httptest.NewRecorder()

	handler.AuthorizePubKey(w, req)
	s.Require().Equal(http.StatusBadRequest, w.Code)
}
