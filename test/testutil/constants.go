package testutil

import (
	"context"
	"errors"
	"net/http"

	"github.com/stretchr/testify/mock"
)

var (
	AnythingOfTypeContext            = mock.MatchedBy(func(_ context.Context) bool { return true })
	AnyThingOfTypeHTTPResponseWriter = mock.MatchedBy(func(_ http.ResponseWriter) bool { return true })
	AnyThingOfTypeHTTPRequest        = mock.MatchedBy(func(_ *http.Request) bool { return true })
	ErrTest                          = errors.New("error")
)

const (
	ValidPubKey            = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIP1ypONGj/C1D/MmsLFNoDUAz7pbqOOkKfJmIqoZMKIv nopain@Nopain-MacBook-Pro.local" //nolint:lll
	ValidPubKeyFingerprint = "8e:16:c3:7b:f3:87:63:48:a9:65:7b:1b:45:ab:c7:90"
	InvalidPubKey          = "ssh-ed25519 XXXX"
)
