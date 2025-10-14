package fnLogger_test

import (
	"fmt"
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

	test.Run("%s", func(t *testing.T) {
		fmt.Printf(
			"%s\n",
			fnLogger.
				NewLogGroup("helloWorldCiaoLee", fnLogger.ColorKeyGreen).String(),
		)

		fmt.Printf(
			"%s\n",
			fnLogger.
				NewLogGroup("hello", fnLogger.ColorKeyGreen).String(),
		)

		fmt.Printf(
			"%s\n",
			fnLogger.
				NewLogGroup("hello_world", fnLogger.ColorKeyGreen).String(),
		)
	})
}
