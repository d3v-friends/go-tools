package fnLogger

import (
	"context"
	"time"

	"github.com/d3v-friends/go-tools/fnCtx"
	"github.com/d3v-friends/go-tools/fnEnv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CtxKeyLogger fnCtx.Key[Logger] = "CTX_LOGGER"
	CtxKeyId     fnCtx.Key[*CtxID] = "CTX_ID"
)

type CtxID struct {
	Id        string
	CreatedAt time.Time
}

func SetID(ctx context.Context, id ...*CtxID) context.Context {
	if len(id) == 1 {
		return fnCtx.Set(ctx, CtxKeyId, id[0])
	}

	return fnCtx.Set(ctx, CtxKeyId, &CtxID{
		Id:        primitive.NewObjectID().Hex(),
		CreatedAt: time.Now(),
	})
}

func GetID(ctx context.Context) (id *CtxID, err error) {
	return fnCtx.Get(ctx, CtxKeyId)
}

func GetIDP(ctx context.Context) *CtxID {
	return fnCtx.GetP(ctx, CtxKeyId)
}

func SetLogger(
	ctx context.Context,
	loggers ...Logger,
) context.Context {
	if len(loggers) == 1 {
		return fnCtx.Set(ctx, CtxKeyLogger, loggers[0])
	}
	return fnCtx.Set(ctx, CtxKeyLogger, NewLogger(LogLevelInfo))
}

func GetLogger(
	ctxes ...context.Context,
) (logger Logger) {
	if len(ctxes) == 1 {
		var err error
		if logger, err = fnCtx.Get(ctxes[0], CtxKeyLogger); err == nil {
			return
		}
	}
	logger = NewLogger(LogLevelInfo)
	return
}

func GetContextLogger(ctx context.Context) (logger ContextLogger) {
	var logLevel = NewLogLevel(fnEnv.String("LOG_LEVEL", "INFO"))
	return NewContextLogger(
		ctx,
		LogGroup{name: "",
			color: ColorKeyWhite,
		},
		logLevel,
	)
}
