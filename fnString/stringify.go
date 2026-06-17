package fnString

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type Args map[string]any

func Stringify(args Args) (res string) {
	if len(args) == 0 {
		return
	}

	var ls = make([]string, len(args))
	var idx = 0
	for k, v := range args {
		ls[idx] = fmt.Sprintf("%s=%s", k, formatValue(v))
		idx += 1
	}

	sort.Strings(ls)

	res = strings.Join(ls, ", ")
	return
}

func formatValue(v any) string {

	switch reflect.TypeOf(v).Kind() {
	case reflect.String:
		return v.(string)
	case reflect.Bool:
		return fmt.Sprintf("%t", v.(bool))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return fmt.Sprintf("%d", v)
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%g", v) // %g는 불필요한 소수점 뒤 0을 제거함
	default:
		return fmt.Sprintf("%v", v)
	}
}
