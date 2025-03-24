package fnLogger_test

import (
	"github.com/d3v-friends/go-tools/fnLogger"
	"testing"
)

func TestLogger(test *testing.T) {
	test.Run("test", func(t *testing.T) {
		var logger = fnLogger.NewLogger()
		logger.Info("hello")
	})
}
