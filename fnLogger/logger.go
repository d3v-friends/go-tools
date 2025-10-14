package fnLogger

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultLogger struct {
	level  LogLevel
	stdout *log.Logger
}

func NewLogger(levels ...LogLevel) Logger {
	var logger = &DefaultLogger{
		stdout: log.Default(),
	}

	if len(levels) == 1 {
		logger.level = levels[0]
	} else {
		logger.level = LogLevelInfo
	}

	logger.stdout.SetOutput(os.Stdout)

	return logger
}

func (x *DefaultLogger) SetLevel(level LogLevel) {
	x.level = level
}

func (x *DefaultLogger) CtxTrace(ctx context.Context, message any, groups ...LogGroup) {
	if x.level < LogLevelTrace {
		return
	}
	x.print(ctx, LogLevelTrace, message, false, groups...)
}

func (x *DefaultLogger) CtxDebug(ctx context.Context, message any, groups ...LogGroup) {
	if x.level < LogLevelDebug {
		return
	}
	x.print(ctx, LogLevelDebug, message, false, groups...)
}

func (x *DefaultLogger) CtxInfo(ctx context.Context, message any, groups ...LogGroup) {
	if x.level < LogLevelInfo {
		return
	}
	x.print(ctx, LogLevelInfo, message, false, groups...)
}

func (x *DefaultLogger) CtxWarn(ctx context.Context, message any, groups ...LogGroup) {
	if x.level < LogLevelWarn {
		return
	}
	x.print(ctx, LogLevelWarn, message, true, groups...)
}

func (x *DefaultLogger) CtxError(ctx context.Context, message any, groups ...LogGroup) {
	if x.level < LogLevelError {
		return
	}
	x.print(ctx, LogLevelError, message, true, groups...)
}

func (x *DefaultLogger) CtxFatal(ctx context.Context, message any, groups ...LogGroup) {
	if x.level < LogLevelFatal {
		return
	}
	x.print(ctx, LogLevelFatal, message, true, groups...)
}

func (x *DefaultLogger) Trace(message any, colors ...LogGroup) {
	x.CtxTrace(context.TODO(), message, colors...)
}

func (x *DefaultLogger) Debug(message any, colors ...LogGroup) {
	x.CtxDebug(context.TODO(), message, colors...)
}

func (x *DefaultLogger) Info(message any, colors ...LogGroup) {
	x.CtxInfo(context.TODO(), message, colors...)
}

func (x *DefaultLogger) Warn(message any, colors ...LogGroup) {
	x.CtxWarn(context.TODO(), message, colors...)
}

func (x *DefaultLogger) Error(message any, colors ...LogGroup) {
	x.CtxError(context.TODO(), message, colors...)
}

func (x *DefaultLogger) Fatal(message any, colors ...LogGroup) {
	x.CtxFatal(context.TODO(), message, colors...)
}

func (x *DefaultLogger) print(
	ctx context.Context,
	level LogLevel,
	message any,
	showLineNumber bool,
	groups ...LogGroup,
) {
	var id, err = GetID(ctx)
	if err != nil {
		id = &CtxID{
			Id:        primitive.NilObjectID.Hex(),
			CreatedAt: time.Now(),
		}
	}

	var group = NilLogGroup
	if len(groups) == 1 {
		group = groups[0]
	}

	var loc, line = x.getLocation()
	var color = ColorKeyGray

	var str = fmt.Sprintf("[%s]%s[%s]\t%s",
		level.log(),
		group.String(),
		id.Id,
		color.log(stringify(message)),
	)

	if showLineNumber {
		str = fmt.Sprintf("%s [%s](%d)", str, loc, line)
	}

	x.stdout.Print(str)
}

func (x *DefaultLogger) getLocation() (loc string, line int) {
	var pc uintptr
	pc, loc, line, _ = runtime.Caller(4)

	var fnName = runtime.FuncForPC(pc).Name()
	var lastSlash = strings.LastIndex(fnName, "/")
	if lastSlash < 0 {
		lastSlash = 0
	}

	var lastDot = strings.LastIndexByte(fnName[lastSlash:], '.') + lastSlash
	loc = fnName[:lastDot]
	return
}
