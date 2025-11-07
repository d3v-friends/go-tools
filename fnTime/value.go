package fnTime

import (
	"time"

	"github.com/d3v-friends/go-tools/fnPanic"
)

var NilTime = fnPanic.Value(time.Parse(time.RFC3339, "2000-01-01T00:00:00Z"))
