package logger

import (
	"testing"

	"github.com/pkg/errors"
)

func TestLoggerOutput(t *testing.T) {
	logger := GetLoggerInstance()
	logger.Debug("Debug message", nil, "ignored value", "key2", "value2")
	logger.Info("Info message", 10.0, "value1", "key2", "value2")
	logger.Warn("", "key1", "value1")
	logger.Error("Error message", errors.New("Raised error"), nil, "ignored", 10.0, "ignored", "key1", "value1", "ignored key")
}
