package fnLogger_test

import (
	"testing"

	"github.com/d3v-friends/go-tools/fnLogger"
)

func TestLogger(test *testing.T) {
	test.Run("test", func(t *testing.T) {
		var logger = fnLogger.NewLogger()
		logger.Info("hello")
	})

	test.Run("level", func(t *testing.T) {
		var logger = fnLogger.NewLogger(fnLogger.LogLevelTrace)
		logger.Trace("trace")
		logger.Debug("debug")
		logger.Info("info")
		logger.Warn("warn")
		logger.Error("error")
		logger.Fatal("fatal")

		logger.SetLevel(fnLogger.LogLevelDebug)
		logger.Trace("trace")
		logger.Debug("debug")
		logger.Info("info")
		logger.Warn("warn")
		logger.Error("error")
		logger.Fatal("fatal")
	})
}
