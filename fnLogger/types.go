package fnLogger

import (
	"context"
	"fmt"
	"regexp"
	"strings"
)

type LogGroup struct {
	name  string
	color ColorKey
}

var NilLogGroup = LogGroup{
	name:  "__________",
	color: ColorKeyGray,
}

// NewLogGroup
// name 은 영문알파벳, 숫자포함 10글자까지 인식한다.
// 영문알파벳은 모두 대문자로 변경된다.
func NewLogGroup(name string, colors ...ColorKey) LogGroup {
	var color = ColorKeyGray
	if len(colors) == 1 {
		color = colors[0]
	}

	var size = 10
	name = regexp.MustCompile("[^a-zA-Z0-9_-]+").ReplaceAllString(name, "")
	if len(name) < size {
		name = name + strings.Repeat("_", size-len(name))
	} else {
		name = name[:size]
	}

	return LogGroup{
		name:  strings.ToUpper(name),
		color: color,
	}
}

func (x LogGroup) String() string {
	return fmt.Sprintf("[%s]", x.color.log(x.name))
}

type Logger interface {
	SetLevel(level LogLevel)
	CtxTrace(ctx context.Context, message any, colors ...LogGroup)
	Trace(message any, colors ...LogGroup)
	CtxDebug(ctx context.Context, message any, colors ...LogGroup)
	Debug(message any, colors ...LogGroup)
	CtxInfo(ctx context.Context, message any, colors ...LogGroup)
	Info(message any, colors ...LogGroup)
	CtxWarn(ctx context.Context, message any, colors ...LogGroup)
	Warn(message any, colors ...LogGroup)
	CtxError(ctx context.Context, message any, colors ...LogGroup)
	Error(message any, colors ...LogGroup)
	CtxFatal(ctx context.Context, message any, colors ...LogGroup)
	Fatal(message any, colors ...LogGroup)
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
		return "INFO_"
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
