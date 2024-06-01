package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func InitializeLogger() *AppLogger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	return &AppLogger{logger: logger}
}

type AppLogger struct {
	logger zerolog.Logger
}

func (l *AppLogger) Debug(msg string) {
	l.logger.Debug().Stack().Msg(msg)
}

func (l *AppLogger) Info(msg string) {
	l.logger.Info().Stack().Msg(msg)
}

func (l *AppLogger) Warn(msg string) {
	l.logger.Warn().Stack().Msg(msg)
}

func (l *AppLogger) Error(msg string, err error) {
	l.logger.Error().Stack().Err(err).Msg(msg)
}

func (l *AppLogger) Fatal(msg string, err error) {
	l.logger.Fatal().Stack().Err(err).Msg(msg)
}

func (l *AppLogger) Panic(msg string, err error) {
	l.logger.Panic().Stack().Err(err).Msg(msg)
}
