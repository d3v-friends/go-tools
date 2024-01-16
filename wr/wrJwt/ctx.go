package wrJwt

import (
	"context"
	"fmt"
)

const CtxJwtKey = "CTX_JWT"

func Set(ctx context.Context, jwt *Jwt) context.Context {
	return context.WithValue(ctx, CtxJwtKey, jwt)
}

func Get(ctx context.Context) (jwt *Jwt, err error) {
	var has bool
	if jwt, has = ctx.Value(CtxJwtKey).(*Jwt); !has {
		err = fmt.Errorf("not found jwt")
		return
	}
	return
}

func GetP(ctx context.Context) (jwt *Jwt) {
	var err error
	if jwt, err = Get(ctx); err != nil {
		return
	}
	return
}
