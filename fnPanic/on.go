package fnPanic

import (
	"fmt"
	"github.com/d3v-friends/go-pure/fnLogger"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

func On(err error) {
	if err != nil {
		panicLogger("%s", err.Error())
		panic(err)
	}
}

func OnPointer[T any](value *T, err error) *T {
	if err != nil {
		panicLogger("%s", err.Error())
		panic(err)
	}
	return value
}

func OnValue[T any](value T, err error) T {
	if err != nil {
		panicLogger("%s", err.Error())
		panic(err)
	}
	return value
}

func Get[T any](value T, err error) T {
	if err != nil {
		panicLogger("%s", err.Error())
		panic(err)
	}
	return value
}

func IsTrue(v bool, err error) {
	if !v {
		panicLogger("%s", err.Error())
		panic(err)
	}
}

func panicLogger(format string, args ...any) {
	var depth = runtime.NumCgoCall()

	var maxSkip = int64(5)
	if depth < 5 {
		maxSkip = depth
	}

	var v = fnLogger.Fields{}
	for i := int64(0); i < maxSkip; i++ {
		var pc, file, line, ok = runtime.Caller(int(i))
		if !ok {
			file = "not_found"
			line = 0
		}

		var fnName = "not_found_fn()"
		var fn = runtime.FuncForPC(pc)
		if fn != nil {
			fnName = fmt.Sprintf("%s()", strings.TrimLeft(filepath.Ext(fn.Name()), "."))
		}

		v[fmt.Sprintf("callstack_%d", i)] = fmt.Sprintf("%s:%d %s", filepath.Base(file), line, fnName)
	}

	v["error"] = fmt.Sprintf(format, args...)
	log.Print(v.ToJson())
}
