package fnGqlgen

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

type CtxInterceptor func(ctx context.Context) context.Context

func Run(
	port string,
	path string,
	handler *handler.Server,
	interceptors ...CtxInterceptor,
) error {
	var interceptor = getInterceptor(interceptors)

	var e = echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTION"},
	}))
	e.POST(path, func(c echo.Context) error {
		var auth = c.Request().Header.Get("Authorization")
		var ctx = interceptor(SetAuth(c.Request().Context(), auth))
		var req = c.Request().WithContext(ctx)
		handler.ServeHTTP(c.Response(), req)
		return nil
	})

	if !strings.HasPrefix(port, ":") {
		port = fmt.Sprintf(":%s", port)
	}

	return e.Start(port)
}

const ctxAuthorization = "CTX_AUTHORIZATION"

func SetAuth(ctx context.Context, auth string) context.Context {
	return context.WithValue(ctx, ctxAuthorization, auth)
}

func GetAuth(ctx context.Context) (auth string, err error) {
	var has bool
	if auth, has = ctx.Value(ctxAuthorization).(string); !has {
		err = fmt.Errorf("not found authorization")
		return
	}
	return
}

func getInterceptor(i []CtxInterceptor) (res CtxInterceptor) {
	res = func(ctx context.Context) context.Context {
		return ctx
	}

	if len(i) != 0 {
		res = i[0]
	}
	return
}
