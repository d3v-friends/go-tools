package fnLogger

import (
	"context"
	"log"
	"os"
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
	printLog(ctx, x.stdout, LogLevelTrace, message, false, groups...)
}

func (x *DefaultLogger) CtxDebug(ctx context.Context, message any, groups ...LogGroup) {
	if x.level < LogLevelDebug {
		return
	}
	printLog(ctx, x.stdout, LogLevelDebug, message, false, groups...)
}

func (x *DefaultLogger) CtxInfo(ctx context.Context, message any, groups ...LogGroup) {
	if x.level < LogLevelInfo {
		return
	}
	printLog(ctx, x.stdout, LogLevelInfo, message, false, groups...)
}

func (x *DefaultLogger) CtxWarn(ctx context.Context, message any, groups ...LogGroup) {
	if x.level < LogLevelWarn {
		return
	}
	printLog(ctx, x.stdout, LogLevelWarn, message, true, groups...)
}

func (x *DefaultLogger) CtxError(ctx context.Context, message any, groups ...LogGroup) {
	if x.level < LogLevelError {
		return
	}
	printLog(ctx, x.stdout, LogLevelError, message, true, groups...)
}

func (x *DefaultLogger) CtxFatal(ctx context.Context, message any, groups ...LogGroup) {
	if x.level < LogLevelFatal {
		return
	}
	printLog(ctx, x.stdout, LogLevelFatal, message, true, groups...)
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
