package fnEnv

import (
	"context"
	"errors"
	"fmt"
)

type EnvMap map[string]string

var ErrNotFoundEnvValue = errors.New("not found env value")

func (x EnvMap) GetString(key fmt.Stringer) (res string) {
	var has bool
	if res, has = x[key.String()]; !has {
		panic(ErrNotFoundEnvValue)
	}
	return
}

/* ------------------------------------------------------------------------------------------------------------ */

type EnvKey string

func (x EnvKey) String() string {
	return string(x)
}

/* ------------------------------------------------------------------------------------------------------------ */

const CtxEnvMap = "CTX_ENV_MAP"

var ErrNotFoundEnvMap = errors.New("not found env map in context")

func SetEnv(ctx context.Context, envMap EnvMap) context.Context {
	return context.WithValue(ctx, CtxEnvMap, envMap)
}

func GetEnv(ctx context.Context) (envMap EnvMap, err error) {
	var has bool
	if envMap, has = ctx.Value(CtxEnvMap).(EnvMap); !has {
		err = ErrNotFoundEnvMap
		return
	}
	return
}

func GetEnvP(ctx context.Context) (envMap EnvMap) {
	var err error
	if envMap, err = GetEnv(ctx); err != nil {
		panic(err)
	}
	return
}
