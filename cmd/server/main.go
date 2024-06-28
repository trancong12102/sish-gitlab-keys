package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/sync/errgroup"

	"github.com/trancong12102/sish-gitlab-keys/internal/config"
	"github.com/trancong12102/sish-gitlab-keys/internal/gitlab"
	"github.com/trancong12102/sish-gitlab-keys/internal/http"
	"github.com/trancong12102/sish-gitlab-keys/internal/telemetry"
)

func main() {
	telemetry.InitLogger()

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

	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	httpServer, err := InitializeHTTPSrv(
		&http.ServerConfig{
			Addr: cfg.ListenAddr,
		},
		&gitlab.ClientConfig{
			URL:         cfg.GitlabURL,
			AccessToken: cfg.GitlabAccessToken,
		},
	)
	if err != nil {
		return fmt.Errorf("initialize http server: %w", err)
	}

	errG.Go(func() error {
		return httpServer.Run()
	})

	errG.Go(func() error {
		<-ctx.Done()

		return httpServer.Shutdown(ctx)
	})

	err = errG.Wait()
	if err != nil {
		return fmt.Errorf("error group wait: %w", err)
	}

	return nil
}
