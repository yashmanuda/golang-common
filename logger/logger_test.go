package logger

import (
	"testing"

	"github.com/pkg/errors"
)

func TestLogger(t *testing.T) {
	logger := GetLoggerInstance()
	logger.Error().Stack().Err(errors.New("Error message")).Msg("Log via initialized logger")
}
