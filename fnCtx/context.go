package fnCtx

import (
	"context"
	"fmt"
	"github.com/d3v-friends/go-pure/fnPanic"
)

func Get[DATA any](ctx context.Context, key string) (res DATA, err error) {
	var isOk bool
	if res, isOk = ctx.Value(key).(DATA); !isOk {
		err = fmt.Errorf("not found value in context: key=%s", key)
		return
	}
	return
}

func GetP[DATA any](ctx context.Context, key string) (res DATA) {
	res = fnPanic.Get(Get[DATA](ctx, key))
	return
}

func Set[DATA any](ctx context.Context, key string, v DATA) context.Context {
	return context.WithValue(ctx, key, v)
}

type FnSetCtx[DATA any] func(ctx context.Context, v DATA) context.Context

func SetFn[DATA any](key string) FnSetCtx[DATA] {
	return func(ctx context.Context, v DATA) context.Context {
		return context.WithValue(ctx, key, v)
	}
}

type FnGetCtx[DATA any] func(ctx context.Context) (DATA, error)

func GetFn[DATA any](key string) FnGetCtx[DATA] {
	return func(ctx context.Context) (res DATA, err error) {
		var isOk bool
		if res, isOk = ctx.Value(key).(DATA); !isOk {
			err = fmt.Errorf("not found value in context: key=%s", key)
			return
		}
		return
	}
}

type FnGetCtxP[DATA any] func(ctx context.Context) DATA

func GetFnP[DATA any](key string) FnGetCtxP[DATA] {
	return func(ctx context.Context) DATA {
		return fnPanic.Get(GetFn[DATA](key)(ctx))
	}
}
