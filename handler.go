package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Authorizer interface {
	AuthorizePubKey(ctx context.Context, key string) error
}

type Handler struct {
	service Authorizer
}

func NewHandler(service Authorizer) *Handler {
	return &Handler{
		service: service,
	}
}

type AuthorizePubKeyRequest struct {
	AuthKey string `json:"auth_key"` //nolint:tagliatelle
}

func (h *Handler) AuthorizePubKey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var reqBody AuthorizePubKeyRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, fmt.Sprintf("decode request body: %v", err), http.StatusBadRequest)

		return
	}

	err = h.service.AuthorizePubKey(ctx, reqBody.AuthKey)
	if err != nil {
		authErr := fmt.Errorf("authorize public key: %w", err)
		w.Header().Set("WWW-Authenticate", authErr.Error())
		http.Error(w, authErr.Error(), http.StatusUnauthorized)

		return
	}

	w.WriteHeader(http.StatusOK)
}
