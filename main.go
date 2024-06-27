package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/sync/errgroup"
)

const (
	ReadHeaderTimeout = 1 * time.Second
)

func main() {
	InitLogger()

	err := run()
	if err != nil {
		slog.Error("run", slog.Any("error", err))
		os.Exit(1)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	errG, ctx := errgroup.WithContext(ctx)

	cfg, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	gitlabClient := NewGitlabClient(cfg.GitlabURL, cfg.GitlabAccessToken)

	service := NewService(gitlabClient)

	handler := NewHandler(service)

	healthcheckHandler, err := NewHealthCheckHandler()
	if err != nil {
		return fmt.Errorf("new healthcheck handler: %w", err)
	}

	router := NewRouter(handler, healthcheckHandler)

	httpServer := &http.Server{
		Addr:              cfg.ListenAddr,
		Handler:           router,
		ReadHeaderTimeout: ReadHeaderTimeout,
	}

	errG.Go(func() error {
		slog.Info("server started", slog.String("addr", cfg.ListenAddr))

		listenErr := httpServer.ListenAndServe()
		if listenErr != nil {
			return fmt.Errorf("http server listen: %w", listenErr)
		}

		return nil
	})

	errG.Go(func() error {
		<-ctx.Done()

		shutdownErr := httpServer.Shutdown(ctx)
		if shutdownErr != nil {
			return fmt.Errorf("http server shutdown: %w", shutdownErr)
		}

		return nil
	})

	err = errG.Wait()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("error group wait: %w", err)
	}

	return nil
}
