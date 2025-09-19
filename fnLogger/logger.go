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

func (x *DefaultLogger) CtxTrace(ctx context.Context, message any, colors ...ColorKey) {
	if x.level < LogLevelTrace {
		return
	}
	x.print(ctx, LogLevelTrace, message, false, colors...)
}

func (x *DefaultLogger) CtxDebug(ctx context.Context, message any, colors ...ColorKey) {
	if x.level < LogLevelDebug {
		return
	}
	x.print(ctx, LogLevelDebug, message, false, colors...)
}

func (x *DefaultLogger) CtxInfo(ctx context.Context, message any, colors ...ColorKey) {
	if x.level < LogLevelInfo {
		return
	}
	x.print(ctx, LogLevelInfo, message, false, colors...)
}

func (x *DefaultLogger) CtxWarn(ctx context.Context, message any, colors ...ColorKey) {
	if x.level < LogLevelWarn {
		return
	}
	x.print(ctx, LogLevelWarn, message, true, colors...)
}

func (x *DefaultLogger) CtxError(ctx context.Context, message any, colors ...ColorKey) {
	if x.level < LogLevelError {
		return
	}
	x.print(ctx, LogLevelError, message, true, colors...)
}

func (x *DefaultLogger) CtxFatal(ctx context.Context, message any, colors ...ColorKey) {
	if x.level < LogLevelFatal {
		return
	}
	x.print(ctx, LogLevelFatal, message, true, colors...)
}

func (x *DefaultLogger) Trace(message any, colors ...ColorKey) {
	x.CtxTrace(context.TODO(), message, colors...)
}

func (x *DefaultLogger) Debug(message any, colors ...ColorKey) {
	x.CtxDebug(context.TODO(), message, colors...)
}

func (x *DefaultLogger) Info(message any, colors ...ColorKey) {
	x.CtxInfo(context.TODO(), message, colors...)
}

func (x *DefaultLogger) Warn(message any, colors ...ColorKey) {
	x.CtxWarn(context.TODO(), message, colors...)
}

func (x *DefaultLogger) Error(message any, colors ...ColorKey) {
	x.CtxError(context.TODO(), message, colors...)
}

func (x *DefaultLogger) Fatal(message any, colors ...ColorKey) {
	x.CtxFatal(context.TODO(), message, colors...)
}

func (x *DefaultLogger) print(
	ctx context.Context,
	level LogLevel,
	message any,
	showLineNumber bool,
	colors ...ColorKey,
) {
	var id, err = GetID(ctx)
	if err != nil {
		id = &CtxID{
			Id:        primitive.NilObjectID.Hex(),
			CreatedAt: time.Now(),
		}
	}

	var loc, line = x.getLocation()
	var color = ColorKeyGray
	if len(colors) == 1 {
		color = colors[0]
	}

	var log = fmt.Sprintf("[%s][%s]\t%s",
		id.Id,
		level.log(),
		color.log(stringify(message)),
	)

	if showLineNumber {
		log = fmt.Sprintf("%s [%s](%d)", log, loc, line)
	}

	x.stdout.Printf(log)
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
