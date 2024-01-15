package fnString

import (
	"fmt"
	"github.com/d3v-friends/go-pure/fnParams"
)

func CommaString(ls []string, iPadding ...bool) (res string) {
	format := "%s,"
	hasPadding := fnParams.Get(iPadding)

	if hasPadding {
		format += " "
	}

	for _, item := range ls {
		res += fmt.Sprintf(format, item)

	}

	if hasPadding {
		res = res[:len(res)-2]
	} else {
		res = res[:len(res)-1]
	}

	return res
}
