package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type Logger interface {

	// don't send error as a part of keysAndValues
	Debug(msg string, keysAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keysAndValues ...interface{})
	
	// error is only mandatory for the below
	Error(msg string, err error, keysAndValues ...interface{})
	Fatal(msg string, err error, keysAndValues ...interface{})
	Panic(msg string, err error, keysAndValues ...interface{})
	
}

type zeroLogger struct {
	logger zerolog.Logger
}

func newZeroLogger() *zeroLogger {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.999Z07:00"
	logger := zerolog.New(os.Stdout).With().Timestamp().Stack().Logger()
	return &zeroLogger{logger: logger}
}

func (z *zeroLogger) logWithKeyValuePairs(level zerolog.Level, msg string, err error, keysAndValues ...interface{}) {
	event := z.logger.WithLevel(level).Stack().Err(err)
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			key, ok := keysAndValues[i].(string)
			if !ok {
				continue
			}
			value := keysAndValues[i+1]
			event = event.Interface(key, value)
		}
	}
	event.Msg(msg)
}

func (z *zeroLogger) Debug(msg string, keysAndValues ...interface{}) {
	z.logWithKeyValuePairs(zerolog.DebugLevel, msg, nil, keysAndValues)
}

func (z *zeroLogger) Info(msg string, keysAndValues ...interface{}) {
	z.logWithKeyValuePairs(zerolog.InfoLevel, msg, nil, keysAndValues...)
}

func (z *zeroLogger) Warn(msg string, keysAndValues ...interface{}) {
	z.logWithKeyValuePairs(zerolog.WarnLevel, msg, nil, keysAndValues...)
}

func (z *zeroLogger) Error(msg string, err error, keysAndValues ...interface{}) {
	z.logWithKeyValuePairs(zerolog.ErrorLevel, msg, err, keysAndValues...)
}

func (z *zeroLogger) Panic(msg string, err error, keysAndValues ...interface{}) {
	z.logWithKeyValuePairs(zerolog.PanicLevel, msg, err, keysAndValues...)
}

func (z *zeroLogger) Fatal(msg string, err error, keysAndValues ...interface{}) {
	z.logWithKeyValuePairs(zerolog.FatalLevel, msg, err, keysAndValues...)
}

var instance Logger
var once sync.Once

func GetLoggerInstance() Logger {
	once.Do(
		func() {
			instance = newZeroLogger()
		})

	return instance
}
