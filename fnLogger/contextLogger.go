package fnLogger

import (
	"context"
	"log"
	"sync"

	"github.com/d3v-friends/go-tools/fnDefault"
)

type DefaultContextLogger struct {
	mutex  *sync.Mutex
	level  LogLevel
	group  LogGroup
	stdout *log.Logger
	ctx    context.Context
}

func (x *DefaultContextLogger) SetLevel(
	level LogLevel,
) ContextLogger {
	x.mutex.Lock()
	x.level = level
	x.mutex.Unlock()
	return x
}

func (x *DefaultContextLogger) SetGroup(
	group LogGroup,
) ContextLogger {
	x.mutex.Lock()
	x.group = group
	x.mutex.Unlock()
	return x
}

func (x *DefaultContextLogger) Trace(message any) {
	if x.level < LogLevelTrace {
		return
	}
	printLog(x.ctx, x.stdout, LogLevelTrace, message, false, x.group)
}

func (x *DefaultContextLogger) Debug(message any) {
	if x.level < LogLevelDebug {
		return
	}
	printLog(x.ctx, x.stdout, LogLevelDebug, message, false, x.group)
}

func (x *DefaultContextLogger) Info(message any) {
	if x.level < LogLevelInfo {
		return
	}
	printLog(x.ctx, x.stdout, LogLevelInfo, message, false, x.group)
}

func (x *DefaultContextLogger) Warn(message any) {
	if x.level < LogLevelWarn {
		return
	}
	printLog(x.ctx, x.stdout, LogLevelWarn, message, true, x.group)
}

func (x *DefaultContextLogger) Error(message any) {
	if x.level < LogLevelError {
		return
	}
	printLog(x.ctx, x.stdout, LogLevelError, message, true, x.group)
}

func (x *DefaultContextLogger) Fatal(message any) {
	if x.level < LogLevelFatal {
		return
	}
	printLog(x.ctx, x.stdout, LogLevelFatal, message, true, x.group)
}

func NewContextLogger(
	ctx context.Context,
	group LogGroup,
	levels ...LogLevel,
) ContextLogger {
	return &DefaultContextLogger{
		mutex:  new(sync.Mutex),
		level:  fnDefault.Param(levels, LogLevelInfo),
		group:  group,
		stdout: log.Default(),
		ctx:    ctx,
	}
}
