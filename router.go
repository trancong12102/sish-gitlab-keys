package main

import (
	"net/http"
)

type AuthHandler interface {
	AuthorizePubKey(w http.ResponseWriter, r *http.Request)
}

type HTTPHandler interface {
	http.Handler
}

func NewRouter(
	authHandler AuthHandler,
	healthcheckHandler HTTPHandler,
) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/health", healthcheckHandler)
	mux.HandleFunc("POST /auth", authHandler.AuthorizePubKey)

	return mux
}
