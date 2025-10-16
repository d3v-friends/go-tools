package fnTime_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/d3v-friends/go-tools/fnTime"
	"github.com/stretchr/testify/assert"
)

func TestTrim(test *testing.T) {
	test.Run("trim", func(t *testing.T) {
		var str = "2025-02-20T01:10:20Z"
		var value, err = time.Parse(time.RFC3339, str)
		assert.NoError(t, err)

		var ymdh = fnTime.TrimYMDH(value)
		assert.Equal(t, "2025-02-20T01:00:00Z", ymdh.Format(time.RFC3339))

		var ymd = fnTime.TrimYMD(value)
		assert.Equal(t, "2025-02-20T00:00:00Z", ymd.Format(time.RFC3339))

		var ym = fnTime.TrimYM(value)
		assert.Equal(t, "2025-02-01T00:00:00Z", ym.Format(time.RFC3339))

		var y = fnTime.TrimY(value)
		assert.Equal(t, "2025-01-01T00:00:00Z", y.Format(time.RFC3339))

		var now = time.Now()
		fmt.Printf("%s\n", fnTime.TrimYMDH(now))
		fmt.Printf("%s\n", fnTime.TrimYMD(now))
	})
}
