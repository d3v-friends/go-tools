package fnJWT

import (
	"context"
	"fmt"
)

const CtxJwtKey = "CTX_JWT"

func SetCtx(ctx context.Context, jwt *Jwt) context.Context {
	return context.WithValue(ctx, CtxJwtKey, jwt)
}

func GetCtx(ctx context.Context) (jwt *Jwt, err error) {
	var has bool
	if jwt, has = ctx.Value(CtxJwtKey).(*Jwt); !has {
		err = fmt.Errorf("not found jwt")
		return
	}
	return
}

func GetCtxP(ctx context.Context) (jwt *Jwt) {
	var err error
	if jwt, err = GetCtx(ctx); err != nil {
		return
	}
	return
}
