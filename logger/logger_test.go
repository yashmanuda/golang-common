package logger

import (
	"errors"
	"testing"
)

func TestLogger(t *testing.T) {
	Logger.Debug("Debug")
	Logger.Info("Info")
	Logger.Warn("Warn")
	Logger.Error("Error without error", nil)
	Logger.Error("Error with error", errors.New("Error message"))
	assertPanic(t, func() { Logger.Panic("Panic without error", nil) })
	assertPanic(t, func() { Logger.Panic("Panic with error", errors.New("Error message")) })
}

func assertPanic(t *testing.T, toExecute func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	toExecute()
}
