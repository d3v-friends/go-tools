package fnCsv

import (
	"fmt"
	"strings"
)

func ToCSV(ls []string) (res string) {
	for _, item := range ls {
		res += fmt.Sprintf("%s,", item)
	}
	res = res[:len(res)-1]
	return res
}

func FromCSV(v string) (ls []string) {
	ls = strings.Split(v, ",")
	for i := range ls {
		ls[i] = strings.Trim(ls[i], " ")
	}
	return
}
