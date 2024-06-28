package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/trancong12102/sish-gitlab-keys/internal/gitlab"
	"github.com/trancong12102/sish-gitlab-keys/internal/ssh"
)

var ErrUserNotActive = errors.New("user not active")

type Service struct {
	pubKeyGetter PubKeyGetter
}

type PubKeyGetter interface {
	GetKeyByFingerprint(ctx context.Context, keyFingerprint string) (*gitlab.Key, error)
}

func NewService(pubKeyGetter PubKeyGetter) *Service {
	return &Service{
		pubKeyGetter: pubKeyGetter,
	}
}

func (s *Service) AuthorizePubKey(ctx context.Context, pubKey string) error {
	fingerprint, err := ssh.GetPubKeyFingerprintLegacyMD5(pubKey)
	if err != nil {
		return fmt.Errorf("get public pubKey fingerprint: %w", err)
	}

	key, err := s.pubKeyGetter.GetKeyByFingerprint(ctx, fingerprint)
	if err != nil {
		return fmt.Errorf("get key by pubKey fingerprint: %w", err)
	}

	if key.User.State != gitlab.UserStateActive {
		return ErrUserNotActive
	}

	return nil
}
