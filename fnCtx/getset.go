// Package fnCtx
// context.Context 에 값을 입출력 할 때 타입캐스팅 등의 중복되는 작업을 줄여주는 함수
package fnCtx

import (
	"context"
	"github.com/d3v-friends/go-tools/fnError"
)

// Key
// Key[매칭할타입] 으로 사용한다.
// example
//
//	func SetSample() {
//		var key = fnCtx.Key[string]("key")
//		var ctx = context.TODO()
//		ctx = fnCtx.Set(ctx, key, "value)
//	}
//
//	func GetSample(ctx context.Context) {
//		var key fnCtx.Key[string] = "key"
//		var value, err = fnCtx.Get(ctx, key)
//	}
type Key[T any] string

func (x Key[T]) String() string {
	return string(x)
}

const (
	ErrNotFoundContextValue = "not_found_context_value"
)

func Get[T any](ctx context.Context, key Key[T]) (value T, err error) {
	var isOk bool
	if value, isOk = ctx.Value(key.String()).(T); !isOk {
		err = fnError.NewFields(ErrNotFoundContextValue, map[string]any{
			"context_key": key.String(),
		})
		return
	}
	return
}

func GetP[T any](ctx context.Context, key Key[T]) (value T) {
	var err error
	if value, err = Get(ctx, key); err != nil {
		panic(err)
	}
	return
}

func Set[T any](ctx context.Context, key Key[T], value T) context.Context {
	return context.WithValue(ctx, key.String(), value)
}
