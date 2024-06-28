package ssh_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/trancong12102/sish-gitlab-keys/internal/ssh"
	"github.com/trancong12102/sish-gitlab-keys/test/testutil"
)

type SSHKeygenTestSuite struct {
	suite.Suite
}

func TestSSHKeygen(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(SSHKeygenTestSuite))
}

func (s *SSHKeygenTestSuite) TestGetPubKeyFingerprintLegacyMD5() {
	fingerprint, err := ssh.GetPubKeyFingerprintLegacyMD5(testutil.ValidPubKey)
	s.Require().NoError(err)
	s.Require().Equal(testutil.ValidPubKeyFingerprint, fingerprint)
}

func (s *SSHKeygenTestSuite) TestGetPubKeyFingerprintLegacyMD5Fail() {
	fingerprint, err := ssh.GetPubKeyFingerprintLegacyMD5(testutil.InvalidPubKey)
	s.Require().Error(err)
	s.Require().Empty(fingerprint)
}
