package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	validPubKey            = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIP1ypONGj/C1D/MmsLFNoDUAz7pbqOOkKfJmIqoZMKIv nopain@Nopain-MacBook-Pro.local" //nolint:lll
	validPubKeyFingerprint = "8e:16:c3:7b:f3:87:63:48:a9:65:7b:1b:45:ab:c7:90"
	invalidPubKey          = "ssh-ed25519 XXXX"
)

type SSHKeygenTestSuite struct {
	suite.Suite
}

func TestSSHKeygen(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(SSHKeygenTestSuite))
}

func (s *SSHKeygenTestSuite) TestGetPubKeyFingerprintLegacyMD5() {
	fingerprint, err := GetPubKeyFingerprintLegacyMD5(validPubKey)
	s.Require().NoError(err)
	s.Require().Equal(validPubKeyFingerprint, fingerprint)
}

func (s *SSHKeygenTestSuite) TestGetPubKeyFingerprintLegacyMD5Fail() {
	fingerprint, err := GetPubKeyFingerprintLegacyMD5(invalidPubKey)
	s.Require().Error(err)
	s.Require().Empty(fingerprint)
}
