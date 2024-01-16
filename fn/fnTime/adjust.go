package fnTime

import (
	"fmt"
	"time"
)

func ToMidnight(t time.Time) (res time.Time, err error) {
	t = t.UTC()
	var str = fmt.Sprintf("%0004d-%02d-%02dT00:00:00Z",
		t.Year(),
		t.Month(),
		t.Day(),
	)

	if res, err = time.Parse(time.RFC3339, str); err != nil {
		return
	}

	return
}
