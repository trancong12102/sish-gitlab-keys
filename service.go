package main

import (
	"context"
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

type Service struct {
	gitlabClient GitlabAPI
}

type GitlabAPI interface {
	GetUserByKeyFingerprint(ctx context.Context, keyFingerprint string) (*User, error)
}

func NewService(gitlabClient GitlabAPI) *Service {
	return &Service{
		gitlabClient: gitlabClient,
	}
}

func (s *Service) AuthorizePubKey(ctx context.Context, key string) error {
	fingerprint, err := GetPubKeyFingerprintLegacyMD5(key)
	if err != nil {
		return fmt.Errorf("get public key fingerprint: %w", err)
	}

	user, err := s.gitlabClient.GetUserByKeyFingerprint(ctx, fingerprint)
	if err != nil {
		return fmt.Errorf("get user by key fingerprint: %w", err)
	}

	if user.ID == nil {
		return ErrUserNotFound
	}

	return nil
}
