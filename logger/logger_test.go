package logger

import (
	"errors"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := InitializeLogger()
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error without error", nil)
	logger.Error("Error with error", errors.New("Error message"))
	assertPanic(t, func() { logger.Panic("Panic without error", nil) })
	assertPanic(t, func() { logger.Panic("Panic with error", errors.New("Error message")) })
}

func assertPanic(t *testing.T, toExecute func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	toExecute()
}
