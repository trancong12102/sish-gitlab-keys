package main

import (
	"context"
	"fmt"

	"github.com/imroc/req/v3"
)

const (
	_apiVersion = "v4"
)

type GitlabClient struct {
	*req.Client
}

func NewGitlabClient(gitlabURL string, accessToken string) *GitlabClient {
	return &GitlabClient{
		Client: req.C().
			SetBaseURL(gitlabURL+"/api/"+_apiVersion).
			SetCommonHeader("Private-Token", accessToken),
	}
}

type User struct {
	ID *int `json:"id,omitempty"`
}

func (c *GitlabClient) GetUserByKeyFingerprint(ctx context.Context, keyFingerprint string) (*User, error) {
	var user User

	err := c.Get("/keys").
		SetQueryParam("fingerprint", keyFingerprint).
		Do(ctx).
		Into(&user)
	if err != nil {
		return nil, fmt.Errorf("get user by key fingerprint: %w", err)
	}

	return &user, nil
}
