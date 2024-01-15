package fnLogger

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type (
	IfLogger interface {
		Trace(format string, args ...any)
		Info(format string, args ...any)
		Warn(format string, args ...any)
		Error(format string, args ...any)
		Fatal(format string, args ...any)
		WithFields(fields ...Fields) IfLogger
		SetLevel(lv Level)
	}
)

type (
	Fields map[string]any
)

func (x *Fields) ToKeyValue() (res string) {
	keyList := make(sort.StringSlice, 0)
	for key := range *x {
		keyList = append(keyList, key)
	}

	keyList.Sort()

	for _, key := range keyList {
		res = fmt.Sprintf("%s=%s\t%s", key, (*x)[key], res)
	}

	res = strings.TrimLeft(res, "\t")

	return
}

func (x *Fields) ToJson() (res string) {
	byteStr, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	return string(byteStr)
}

type (
	Level int
)

const (
	Trace Level = 0
	Info  Level = 1
	Warn  Level = 2
	Error Level = 3
	Fatal Level = 4
)

func (x *Level) String() string {
	switch *x {
	case 0:
		return "Trace"
	case 1:
		return "Info"
	case 2:
		return "Warn"
	case 3:
		return "Error"
	case 4:
		return "Fatal"
	default:
		panic(fmt.Errorf("unknown logger level: level=%d", *x))
	}
}

func MergeFields(fields ...Fields) (res Fields) {
	res = make(Fields)
	for _, field := range fields {
		for key, stringer := range field {
			res[key] = stringer
		}
	}
	return
}
