package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var Logger = initializeLogger()

type appLogger struct {
	logger zerolog.Logger
}

func initializeLogger() *appLogger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	return &appLogger{logger: logger}
}

func (l *appLogger) Debug(msg string) {
	l.logger.Debug().Stack().Msg(msg)
}

func (l *appLogger) Info(msg string) {
	l.logger.Info().Stack().Msg(msg)
}

func (l *appLogger) Warn(msg string) {
	l.logger.Warn().Stack().Msg(msg)
}

func (l *appLogger) Error(msg string, err error) {
	l.logger.Error().Stack().Err(err).Msg(msg)
}

func (l *appLogger) Fatal(msg string, err error) {
	l.logger.Fatal().Stack().Err(err).Msg(msg)
}

func (l *appLogger) Panic(msg string, err error) {
	l.logger.Panic().Stack().Err(err).Msg(msg)
}
