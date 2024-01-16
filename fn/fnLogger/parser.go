package fnLogger

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

type (
	DefaultPrinter struct{}
)

func (x *DefaultPrinter) Print(v Fields) {
	var pc, file, line, ok = runtime.Caller(3)
	if !ok {
		file = "not_found"
		line = 0
	}

	var fnName = "not_found_fn()"
	var fn = runtime.FuncForPC(pc)
	if fn != nil {
		fnName = fmt.Sprintf("%s()", strings.TrimLeft(filepath.Ext(fn.Name()), "."))
	}

	v["code"] = fmt.Sprintf("%s:%d %s", filepath.Base(file), line, fnName)

	log.Print(v.ToJson())
}

func NewDefaultLogger() IfLogger {
	return NewLogger(&DefaultPrinter{})
}

/* ------------------------------------------------------------------------------------------------------------ */

// json parser

func ToJson(v any) (str string, err error) {
	var byteStr []byte
	if byteStr, err = json.Marshal(v); err != nil {
		return
	}
	str = string(byteStr)
	return
}

func ToJsonP(v any) (str string) {
	var err error
	if str, err = ToJson(v); err != nil {
		panic(err)
	}
	return
}

//
