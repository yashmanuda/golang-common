package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

// singleton instance
var instance zerolog.Logger
var once sync.Once

func GetLoggerInstance() zerolog.Logger {
	once.Do(
		func() {
			zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
			instance = zerolog.New(os.Stderr).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{
				Out:        os.Stderr,
				TimeFormat: "2006-01-02T15:04:05.999Z07:00",
			})
		})

	return instance
}
