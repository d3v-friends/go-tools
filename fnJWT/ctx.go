package fnJWT

import (
	"context"
	"fmt"
)

const ctxJwt = "CTX_JWT"

func SetJWT(ctx context.Context, jwt *JWT) context.Context {
	return context.WithValue(ctx, ctxJwt, jwt)
}

func GetJWT(ctx context.Context) (jwt *JWT, err error) {
	var has bool
	if jwt, has = ctx.Value(ctxJwt).(*JWT); !has {
		err = fmt.Errorf("not found jwt")
		return
	}
	return
}

func GetJWTP(ctx context.Context) (jwt *JWT) {
	var err error
	if jwt, err = GetJWT(ctx); err != nil {
		return
	}
	return
}
