package fnLogger

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

func Panic(format string, args ...any) {
	var err = fmt.Errorf(format, args...)
	panicLogger(err)
}

func panicLogger(err error) {
	if err == nil {
		return
	}

	var depth = runtime.NumCgoCall()

	var maxSkip = int64(5)
	if depth < 5 {
		maxSkip = depth
	}

	var v = Fields{}
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

	v["error"] = err.Error()
	log.Panicln(v.ToJson())
}
