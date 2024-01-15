package fnContext

import (
	"context"
	"fmt"
)

func Set(ctx context.Context, key string, value any) context.Context {
	return context.WithValue(ctx, key, value)
}

func Get[T any](ctx context.Context, key string) (T, error) {
	v, isOk := ctx.Value(key).(T)
	if !isOk {
		return *new(T), fmt.Errorf("not found context key: key=%s", key)
	}
	return v, nil
}
