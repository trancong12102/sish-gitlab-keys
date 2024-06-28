package telemetry_test

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	slogzerolog "github.com/samber/slog-zerolog/v2"
	"github.com/stretchr/testify/suite"

	"github.com/trancong12102/sish-gitlab-keys/internal/telemetry"
)

type LoggerTestSuite struct {
	suite.Suite
}

func TestLogger(t *testing.T) { //nolint:paralleltest
	suite.Run(t, new(LoggerTestSuite))
}

func (s *LoggerTestSuite) TestInitLogger_Production() {
	s.T().Setenv("APP_ENV", "production")
	telemetry.InitLogger()
	s.Require().NotNil(slog.Default())

	zerologLogger := zerolog.New(os.Stdout).With().Logger()
	zerologHandler := slogzerolog.Option{
		Level:  slog.LevelDebug,
		Logger: &zerologLogger,
	}.NewZerologHandler()
	logger := slog.New(zerologHandler)
	s.Require().Equal(logger, slog.Default())
}

func (s *LoggerTestSuite) TestInitLogger_Development() {
	s.T().Setenv("APP_ENV", "development")
	telemetry.InitLogger()
	s.Require().NotNil(slog.Default())

	zerologLogger := zerolog.New(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		},
	).With().Logger()
	zerologHandler := slogzerolog.Option{
		Level:  slog.LevelDebug,
		Logger: &zerologLogger,
	}.NewZerologHandler()
	logger := slog.New(zerologHandler)
	s.Require().Equal(logger, slog.Default())
}
