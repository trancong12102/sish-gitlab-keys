package main

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
)

var (
	anythingOfTypeContext            = mock.MatchedBy(func(_ context.Context) bool { return true })
	anyThingOfTypeHTTPResponseWriter = mock.MatchedBy(func(_ http.ResponseWriter) bool { return true })
	anyThingOfTypeHTTPRequest        = mock.MatchedBy(func(_ *http.Request) bool { return true })
	errTest                          = errors.New("error")
	gitlabUserID                     = 1
	gitlabUser                       = &User{ID: &gitlabUserID}
)

func TestMain(m *testing.M) {
	code := m.Run()
	if code != 0 {
		panic("Tests failed")
	}
}
