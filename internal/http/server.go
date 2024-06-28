package http

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

const (
	_readHeaderTimeout = 1 * time.Second
)

type Server struct {
	*http.Server
}

type ServerHandler interface {
	http.Handler
}

type ServerConfig struct {
	Addr string
}

func NewHTTPSrv(
	handler ServerHandler,
	config *ServerConfig,
) *Server {
	return &Server{
		Server: &http.Server{
			Addr:              config.Addr,
			Handler:           handler,
			ReadHeaderTimeout: _readHeaderTimeout,
		},
	}
}

func (s *Server) Run() error {
	slog.Info("server started", slog.String("addr", s.Addr))

	if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve http server: %w", err)
	}

	return nil
}
