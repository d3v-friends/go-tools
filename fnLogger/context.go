package fnLogger

import (
	"context"
	"fmt"
)

const CtxLogger = "CTX_LOGGER"

func Get(ctx context.Context, iDefault ...IfLogger) (logger IfLogger) {
	var isOk bool
	if logger, isOk = ctx.Value(CtxLogger).(IfLogger); !isOk {
		if len(iDefault) == 0 {
			panic(fmt.Errorf("not found logger"))
		}
		return iDefault[0]
	}
	return
}

func Set(ctx context.Context, logger IfLogger) context.Context {
	return context.WithValue(ctx, CtxLogger, logger)
}
