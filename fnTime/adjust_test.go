package fnTime

import (
	"fmt"
	"github.com/d3v-friends/go-pure/fnPanic"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAdjust(test *testing.T) {
	test.Run("test", func(t *testing.T) {
		var now = time.Now()

		var nowStr = now.Format(time.RFC3339)
		var utc = now.UTC()
		var resStr = fmt.Sprintf("%0004d-%02d-%02dT00:00:00Z", utc.Year(), utc.Month(), utc.Day())

		var timeNow = fnPanic.OnValue(time.Parse(time.RFC3339, nowStr))

		var parsed = fnPanic.OnValue(ToMidnight(timeNow))

		assert.Equal(t, resStr, parsed.Format(time.RFC3339))

	})
}
