package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/rs/zerolog"
	slogzerolog "github.com/samber/slog-zerolog/v2"
)

func InitLogger() {
	var zerologLogger zerolog.Logger

	if os.Getenv("APP_ENV") == string(EnvironmentTypeProduction) {
		zerologLogger = zerolog.New(os.Stdout).With().Logger()
	} else {
		zerologLogger = zerolog.New(
			zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: time.RFC3339,
			},
		).With().Logger()
	}

	zerologHandler := slogzerolog.Option{
		Level:  slog.LevelDebug,
		Logger: &zerologLogger,
	}.NewZerologHandler()

	logger := slog.New(zerologHandler)

	slog.SetDefault(logger)
}
