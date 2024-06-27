package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
}

func TestService(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ServiceTestSuite))
}

func (s *ServiceTestSuite) TestAuthorizePubKey() {
	gitlabClient := NewMockGitlabAPI(s.T())

	gitlabClient.EXPECT().
		GetUserByKeyFingerprint(anythingOfTypeContext, validPubKeyFingerprint).
		Once().
		Return(gitlabUser, nil)

	service := NewService(gitlabClient)
	err := service.AuthorizePubKey(context.Background(), validPubKey)
	s.Require().NoError(err)
}

func (s *ServiceTestSuite) TestAuthorizePubKey_InvalidKey() {
	service := NewService(nil)
	err := service.AuthorizePubKey(context.Background(), invalidPubKey)
	s.Require().Error(err)
	s.Require().Equal("get public key fingerprint: parse authorized key: ssh: no key found", err.Error())
}

func (s *ServiceTestSuite) TestAuthorizePubKey_UserNotFound() {
	gitlabClient := NewMockGitlabAPI(s.T())

	gitlabClient.EXPECT().
		GetUserByKeyFingerprint(anythingOfTypeContext, validPubKeyFingerprint).
		Once().
		Return(&User{}, nil)

	service := NewService(gitlabClient)
	err := service.AuthorizePubKey(context.Background(), validPubKey)
	s.Require().Error(err)
	s.Require().Equal(ErrUserNotFound, err)
}

func (s *ServiceTestSuite) TestAuthorizePubKey_GetUserError() {
	gitlabClient := NewMockGitlabAPI(s.T())

	gitlabClient.EXPECT().
		GetUserByKeyFingerprint(anythingOfTypeContext, validPubKeyFingerprint).
		Once().
		Return(nil, errTest)

	service := NewService(gitlabClient)
	err := service.AuthorizePubKey(context.Background(), validPubKey)
	s.Require().Error(err)
	s.Require().Equal("get user by key fingerprint: error", err.Error())
}
