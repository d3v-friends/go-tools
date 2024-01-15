package fnPanic

import (
	"fmt"
	"testing"
)

func TestPanic(test *testing.T) {
	test.Run("on", func(t *testing.T) {
		var fn = func() (err error) {
			err = fmt.Errorf("on error")
			return
		}

		On(fn())

	})
}
