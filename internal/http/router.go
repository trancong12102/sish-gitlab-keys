package http

import (
	"net/http"
)

type AuthHandler interface {
	AuthorizePubKey(w http.ResponseWriter, r *http.Request)
}

type RouterHandler interface {
	http.Handler
}

var _ http.Handler = (*Router)(nil)

type Router struct {
	handler http.Handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.handler.ServeHTTP(w, req)
}

func NewRouter(
	authHandler AuthHandler,
	healthcheck RouterHandler,
) *Router {
	mux := http.NewServeMux()
	mux.Handle("/health", healthcheck)
	mux.HandleFunc("POST /auth", authHandler.AuthorizePubKey)

	return &Router{
		handler: mux,
	}
}
