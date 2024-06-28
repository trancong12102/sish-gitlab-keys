package auth_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/trancong12102/sish-gitlab-keys/internal/auth"
	"github.com/trancong12102/sish-gitlab-keys/internal/gitlab"
	authmocks "github.com/trancong12102/sish-gitlab-keys/mocks/auth"
	"github.com/trancong12102/sish-gitlab-keys/test/testutil"
)

type ServiceTestSuite struct {
	suite.Suite
}

func TestService(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ServiceTestSuite))
}

func (s *ServiceTestSuite) TestAuthorizePubKey() {
	gitlabClient := authmocks.NewMockPubKeyGetter(s.T())

	gitlabClient.EXPECT().
		GetKeyByFingerprint(testutil.AnythingOfTypeContext, testutil.ValidPubKeyFingerprint).
		Once().
		Return(&gitlab.Key{
			ID: 1,
			User: gitlab.User{
				ID:    1,
				State: gitlab.UserStateActive,
			},
		}, nil)

	service := auth.NewService(gitlabClient)
	err := service.AuthorizePubKey(context.Background(), testutil.ValidPubKey)
	s.Require().NoError(err)
}

func (s *ServiceTestSuite) TestAuthorizePubKey_InvalidKey() {
	service := auth.NewService(nil)
	err := service.AuthorizePubKey(context.Background(), testutil.InvalidPubKey)
	s.Require().Error(err)
	s.Require().Equal("get public pubKey fingerprint: parse authorized key: ssh: no key found", err.Error())
}

func (s *ServiceTestSuite) TestAuthorizePubKey_GetKeyError() {
	gitlabClient := authmocks.NewMockPubKeyGetter(s.T())

	gitlabClient.EXPECT().
		GetKeyByFingerprint(testutil.AnythingOfTypeContext, testutil.ValidPubKeyFingerprint).
		Once().
		Return(nil, testutil.ErrTest)

	service := auth.NewService(gitlabClient)
	err := service.AuthorizePubKey(context.Background(), testutil.ValidPubKey)
	s.Require().Error(err)
	s.Require().Equal("get key by pubKey fingerprint: error", err.Error())
}

func (s *ServiceTestSuite) TestAuthorizePubKey_UserNotActive() {
	gitlabClient := authmocks.NewMockPubKeyGetter(s.T())

	gitlabClient.EXPECT().
		GetKeyByFingerprint(testutil.AnythingOfTypeContext, testutil.ValidPubKeyFingerprint).
		Once().
		Return(&gitlab.Key{
			ID: 1,
			User: gitlab.User{
				ID:    1,
				State: "blocked",
			},
		}, nil)

	service := auth.NewService(gitlabClient)
	err := service.AuthorizePubKey(context.Background(), testutil.ValidPubKey)
	s.Require().Equal(auth.ErrUserNotActive, err)
}
