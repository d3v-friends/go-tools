package fnLogger

import (
	"fmt"
	"time"
)

type (
	Logger struct {
		// values
		level  Level
		fields Fields

		// fns
		printer IfPrinter
	}

	IfPrinter interface {
		Print(v Fields)
	}
)

func (x *Logger) SetLevel(lv Level) {
	x.level = lv
}

func (x *Logger) Trace(format string, args ...any) {
	if Trace < x.level {
		return
	}
	x.print(format, args...)
}

func (x *Logger) Info(format string, args ...any) {
	if Info < x.level {
		return
	}
	x.print(format, args...)
}

func (x *Logger) Warn(format string, args ...any) {
	if Warn < x.level {
		return
	}
	x.print(format, args...)
}

func (x *Logger) Error(format string, args ...any) {
	if Error < x.level {
		return
	}
	x.print(format, args...)
}

func (x *Logger) Fatal(format string, args ...any) {
	if Fatal < x.level {
		return
	}
	x.print(format, args...)
}

func (x *Logger) WithFields(fields ...Fields) IfLogger {
	ls := make([]Fields, 0)
	ls = append(ls, x.fields)
	ls = append(ls, fields...)

	logger := NewLogger(x.printer, ls...)
	logger.SetLevel(x.level)
	return logger
}

func (x *Logger) print(format string, args ...any) {
	res := make(Fields)
	for key, value := range x.fields {
		res[key] = value
	}
	res["level"] = x.level.String()
	res["time"] = time.Now().Format(time.RFC3339)
	res["log"] = fmt.Sprintf(format, args...)

	x.printer.Print(res)
}

func NewLogger(
	printer IfPrinter,
	fields ...Fields,
) IfLogger {
	return &Logger{
		level:   Info,
		fields:  MergeFields(fields...),
		printer: printer,
	}
}
