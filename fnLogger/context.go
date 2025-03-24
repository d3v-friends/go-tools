package fnLogger

import (
	"context"
	"github.com/d3v-friends/go-tools/fnCtx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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
