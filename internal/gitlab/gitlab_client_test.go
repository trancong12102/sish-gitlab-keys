package gitlab_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"

	"github.com/trancong12102/sish-gitlab-keys/internal/gitlab"
	"github.com/trancong12102/sish-gitlab-keys/test/testutil"
)

const (
	_gitlabURL         = "https://example.com"
	_gitlabAPIBaseURL  = _gitlabURL + "/api/v4"
	_gitlabAccessToken = "token"
)

var _gitlabClientConfig = &gitlab.ClientConfig{
	URL:         _gitlabURL,
	AccessToken: _gitlabAccessToken,
}

type GitlabClientTestSuite struct {
	suite.Suite
}

func TestGitlabClient(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(GitlabClientTestSuite))
}

func (s *GitlabClientTestSuite) TestNewGitlabClient() {
	client := gitlab.NewClient(_gitlabClientConfig)
	s.Require().NotNil(client)
	s.Require().Equal(_gitlabAPIBaseURL, client.BaseURL)
	s.Require().Equal(_gitlabAccessToken, client.Headers.Get("Private-Token"))
	s.Require().Equal("application/json", client.Headers.Get("Accept"))
}

func (s *GitlabClientTestSuite) TestGetKeyByFingerprint() {
	client := gitlab.NewClient(_gitlabClientConfig)

	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	responder := httpmock.NewStringResponder(200, `{"id":1,"user":{"id":1,"state":"active"}}`)

	endpoint := _gitlabAPIBaseURL + "/keys?fingerprint=" + url.QueryEscape(testutil.ValidPubKeyFingerprint)
	httpmock.RegisterResponder("GET", endpoint, responder)

	ctx := context.Background()
	key, err := client.GetKeyByFingerprint(ctx, testutil.ValidPubKeyFingerprint)
	s.Require().NoError(err)
	s.Require().Equal(1, key.ID)
	s.Require().Equal(1, key.User.ID)
	s.Require().Equal(gitlab.UserStateActive, key.User.State)
}

func (s *GitlabClientTestSuite) TestGetKeyByFingerprint_KeyNotFound() {
	client := gitlab.NewClient(_gitlabClientConfig)

	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	responder := httpmock.NewStringResponder(200, "{}")

	endpoint := _gitlabAPIBaseURL + "/keys?fingerprint=" + url.QueryEscape(testutil.ValidPubKeyFingerprint)
	httpmock.RegisterResponder("GET", endpoint, responder)

	ctx := context.Background()
	key, err := client.GetKeyByFingerprint(ctx, testutil.ValidPubKeyFingerprint)
	s.Require().Error(err)
	s.Require().Equal(gitlab.ErrKeyNotFound.Error(), err.Error())
	s.Require().Nil(key)
}

func (s *GitlabClientTestSuite) TestGetKeyByFingerprint_Error() {
	client := gitlab.NewClient(_gitlabClientConfig)

	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	responder := httpmock.NewErrorResponder(testutil.ErrTest)

	endpoint := _gitlabAPIBaseURL + "/keys?fingerprint=" + url.QueryEscape(testutil.ValidPubKeyFingerprint)
	httpmock.RegisterResponder("GET", endpoint, responder)

	ctx := context.Background()
	key, err := client.GetKeyByFingerprint(ctx, testutil.ValidPubKeyFingerprint)
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "get key by fingerprint: ")
	s.Require().Nil(key)
}

func (s *GitlabClientTestSuite) TestGetKeyByFingerprint_StatusNotOK() {
	client := gitlab.NewClient(_gitlabClientConfig)

	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	responder := httpmock.NewStringResponder(404, "")

	endpoint := _gitlabAPIBaseURL + "/keys?fingerprint=" + url.QueryEscape(testutil.ValidPubKeyFingerprint)
	httpmock.RegisterResponder("GET", endpoint, responder)

	ctx := context.Background()
	key, err := client.GetKeyByFingerprint(ctx, testutil.ValidPubKeyFingerprint)
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "get key by fingerprint: ")
	s.Require().Nil(key)
}
