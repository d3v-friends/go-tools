package fnJWT

import (
	"context"
	"fmt"
)

const ctxJwt = "CTX_JWT"

func SetJWT[T IfJwtData](ctx context.Context, jwt *JWT[T]) context.Context {
	return context.WithValue(ctx, ctxJwt, jwt)
}

func GetJWT[T IfJwtData](ctx context.Context) (jwt *JWT[T], err error) {
	var has bool
	if jwt, has = ctx.Value(ctxJwt).(*JWT[T]); !has {
		err = fmt.Errorf("not found jwt")
		return
	}
	return
}

func GetJWTP[T IfJwtData](ctx context.Context) (jwt *JWT[T]) {
	var err error
	if jwt, err = GetJWT[T](ctx); err != nil {
		return
	}
	return
}
