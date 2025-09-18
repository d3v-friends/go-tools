package fnLogger

import (
	"context"
	"fmt"
	"strings"
)

type Logger interface {
	SetLevel(level LogLevel)
	CtxTrace(ctx context.Context, message any, colors ...ColorKey)
	Trace(message any, colors ...ColorKey)
	CtxDebug(ctx context.Context, message any, colors ...ColorKey)
	Debug(message any, colors ...ColorKey)
	CtxInfo(ctx context.Context, message any, colors ...ColorKey)
	Info(message any, colors ...ColorKey)
	CtxWarn(ctx context.Context, message any, colors ...ColorKey)
	Warn(message any, colors ...ColorKey)
	CtxError(ctx context.Context, message any, colors ...ColorKey)
	Error(message any, colors ...ColorKey)
	CtxFatal(ctx context.Context, message any, colors ...ColorKey)
	Fatal(message any, colors ...ColorKey)
}

type ColorKey string

func (x ColorKey) String() string {
	return string(x)
}

func (x ColorKey) log(v string) string {
	return fmt.Sprintf("%s%s\u001B[0m", x.String(), v)
}

const (
	ColorKeyRed     ColorKey = "\033[31m"
	ColorKeyGreen   ColorKey = "\033[32m"
	ColorKeyYellow  ColorKey = "\033[33m"
	ColorKeyBlue    ColorKey = "\033[34m"
	ColorKeyMagenta ColorKey = "\033[35m"
	ColorKeyCyan    ColorKey = "\033[36m"
	ColorKeyGray    ColorKey = "\033[37m"
	ColorKeyWhite   ColorKey = "\033[97m"
	ColorKeyReset   ColorKey = "\033[0m"
)

type LogLevel int

const (
	LogLevelInvalid LogLevel = iota
	LogLevelFatal
	LogLevelError
	LogLevelWarn
	LogLevelInfo
	LogLevelDebug
	LogLevelTrace
)

var LogLevelAll = []LogLevel{
	LogLevelTrace,
	LogLevelDebug,
	LogLevelInfo,
	LogLevelWarn,
	LogLevelError,
	LogLevelFatal,
}

func NewLogLevel(str string, defaults ...LogLevel) LogLevel {
	switch strings.ToLower(str) {
	case "fatal":
		return LogLevelFatal
	case "error":
		return LogLevelError
	case "warn":
		return LogLevelWarn
	case "info":
		return LogLevelInfo
	case "debug":
		return LogLevelDebug
	case "trace":
		return LogLevelTrace
	default:
		if len(defaults) == 1 {
			return defaults[0]
		}
		return LogLevelInfo
	}
}

func (x LogLevel) Int() int {
	return int(x)
}

func (x LogLevel) Valid() bool {
	for _, level := range LogLevelAll {
		if x == level {
			return true
		}
	}
	return false
}

func (x LogLevel) String() string {
	switch x {
	case LogLevelTrace:
		return "TRACE"
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO_"
	case LogLevelWarn:
		return "WARN_"
	case LogLevelError:
		return "ERROR"
	case LogLevelFatal:
		return "FATAL"
	default:
		return "INFO"
	}
}

func (x LogLevel) log() string {
	switch x {
	case LogLevelTrace:
		return ColorKeyGray.log(x.String())
	case LogLevelDebug:
		return ColorKeyWhite.log(x.String())
	case LogLevelInfo:
		return ColorKeyGreen.log(x.String())
	case LogLevelWarn:
		return ColorKeyYellow.log(x.String())
	case LogLevelError:
		return ColorKeyMagenta.log(x.String())
	case LogLevelFatal:
		return ColorKeyRed.log(x.String())
	default:
		return ColorKeyGray.log(x.String())
	}
}
