package auth_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/trancong12102/sish-gitlab-keys/internal/auth"
	authmocks "github.com/trancong12102/sish-gitlab-keys/mocks/auth"
	"github.com/trancong12102/sish-gitlab-keys/test/testutil"
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
	validBody        = []byte(`{"auth_key":"` + testutil.ValidPubKey + `"}`)
	notFoundUserKey  = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIP1ypONGj/C1D/MmsLFNoDUAz7pbqOOkKfJmIqoZMKIv nopain@Nopain-MacBook-Pro.local" //nolint:lll
	notFoundUserBody = []byte(`{"auth_key":"` + notFoundUserKey + `"}`)
	invalidBody      = []byte("xxx")
)

func (s *HandlerTestSuite) TestAuthorizePubKey() {
	service := authmocks.NewMockAuthorizer(s.T())

	service.EXPECT().
		AuthorizePubKey(testutil.AnythingOfTypeContext, testutil.ValidPubKey).
		Once().
		Return(nil)

	handler := auth.NewHandler(service)

	body := bytes.NewReader(validBody)
	req := httptest.NewRequest(http.MethodGet, reqURL, body)

	w := httptest.NewRecorder()

	handler.AuthorizePubKey(w, req)
	s.Require().Equal(http.StatusOK, w.Code)
}

func (s *HandlerTestSuite) TestAuthorizePubKey_UserNotActive() {
	service := authmocks.NewMockAuthorizer(s.T())

	service.EXPECT().
		AuthorizePubKey(testutil.AnythingOfTypeContext, notFoundUserKey).
		Once().
		Return(auth.ErrUserNotActive)

	handler := auth.NewHandler(service)

	body := bytes.NewReader(notFoundUserBody)
	req := httptest.NewRequest(http.MethodGet, reqURL, body)

	w := httptest.NewRecorder()

	handler.AuthorizePubKey(w, req)
	s.Require().Equal(http.StatusUnauthorized, w.Code)
}

func (s *HandlerTestSuite) TestAuthorizePubKey_InvalidBody() {
	service := authmocks.NewMockAuthorizer(s.T())

	handler := auth.NewHandler(service)

	body := bytes.NewReader(invalidBody)
	req := httptest.NewRequest(http.MethodGet, reqURL, body)

	w := httptest.NewRecorder()

	handler.AuthorizePubKey(w, req)
	s.Require().Equal(http.StatusBadRequest, w.Code)
}
