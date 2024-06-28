package gitlab

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/imroc/req/v3"
)

const (
	_gitlabAPIVersion    = "v4"
	_gitlabClientTimeout = 1 * time.Minute
	_gitlabClientRetry   = 3
)

var ErrKeyNotFound = errors.New("key not found")

type Client struct {
	*req.Client
}

type ClientConfig struct {
	URL         string
	AccessToken string
}

func NewClient(config *ClientConfig) *Client {
	return &Client{
		Client: req.C().
			SetBaseURL(config.URL+"/api/"+_gitlabAPIVersion).
			SetCommonHeader("Private-Token", config.AccessToken).
			SetCommonHeader("Accept", "application/json").
			SetTimeout(_gitlabClientTimeout).
			SetCommonRetryCount(_gitlabClientRetry),
	}
}

const (
	UserStateActive = "active"
)

type User struct {
	ID    int    `json:"id,omitempty"`
	State string `json:"state,omitempty"`
}

type Key struct {
	ID   int  `json:"id,omitempty"`
	User User `json:"user,omitempty"`
}

func (c *Client) GetKeyByFingerprint(ctx context.Context, keyFingerprint string) (*Key, error) {
	var key Key

	err := c.Get("/keys").
		SetQueryParam("fingerprint", keyFingerprint).
		Do(ctx).
		Into(&key)
	if err != nil {
		return nil, fmt.Errorf("get key by fingerprint: %w", err)
	}

	if key == (Key{}) {
		return nil, ErrKeyNotFound
	}

	return &key, nil
}
