package fnCtx

import "context"

func Get[T any](ctx context.Context, key string, noValueErr error) (res T, err error) {
	var isOk bool
	if res, isOk = ctx.Value(key).(T); !isOk {
		err = noValueErr
		return
	}
	return
}

func GetP[T any](ctx context.Context, key string, noValueError error) (res T) {
	var err error
	if res, err = Get[T](ctx, key, noValueError); err != nil {
		panic(err)
	}
	return
}

func Set[T any](ctx context.Context, key string, value T) context.Context {
	return context.WithValue(ctx, key, value)
}
