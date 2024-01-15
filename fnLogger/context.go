package fnLogger

import (
	"context"
	"fmt"
	"github.com/d3v-friends/go-pure/fnVars"
)

func Get(ctx context.Context, iDefault ...IfLogger) (logger IfLogger) {
	var isOk bool
	if logger, isOk = ctx.Value(fnVars.CTX_LOGGER).(IfLogger); !isOk {
		if len(iDefault) == 0 {
			panic(fmt.Errorf("not found logger"))
		}
		return iDefault[0]
	}
	return
}

func Set(ctx context.Context, logger IfLogger) context.Context {
	return context.WithValue(ctx, fnVars.CTX_LOGGER, logger)
}
