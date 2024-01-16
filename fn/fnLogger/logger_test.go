package fnLogger

import (
	"testing"
	"time"
)

func TestLogger(test *testing.T) {
	var logger = NewDefaultLogger()
	logger.SetLevel(Trace)

	test.Run("run", func(t *testing.T) {
		logger.
			WithFields(Fields{
				"time": time.Now().Format(time.RFC3339),
			}).
			Info("hello")

	})
}
