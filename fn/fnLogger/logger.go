package fnLogger

import (
	"context"
	"errors"
	"github.com/d3v-friends/go-tools/fn/fnCtx"
	"log"
)

const CtxLogger = "CTX_LOGGER"

var (
	ErrNotFoundLoggerInContext = errors.New("not found logger in context")
)

func NewLogger(requestIds ...string) (res *log.Logger) {
	res = log.Default()
	res.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	if len(requestIds) != 0 {
		res.SetPrefix(requestIds[0])
	}
	return
}

func Set(ctx context.Context, logger *log.Logger) context.Context {
	return context.WithValue(ctx, CtxLogger, logger)
}

func Get(ctx context.Context) (res *log.Logger, err error) {
	return fnCtx.Get[*log.Logger](ctx, CtxLogger, ErrNotFoundLoggerInContext)
}

func GetP(ctx context.Context) (res *log.Logger) {
	return fnCtx.GetP[*log.Logger](ctx, CtxLogger, ErrNotFoundLoggerInContext)
}
