package wrGqlgen

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
)

type CtxInterceptor func(ctx context.Context) context.Context

func Run(
	port string,
	path string,
	handler *handler.Server,
	hasPlayground bool,
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
		var ctx = c.Request().Context()
		ctx = SetHeader(ctx, c.Request().Header)
		ctx = interceptor(ctx)

		var req = c.Request().WithContext(ctx)
		handler.ServeHTTP(c.Response(), req)
		return nil
	})

	if hasPlayground {
		e.GET("/playground", func(c echo.Context) error {
			playground.ApolloSandboxHandler("playground", path).ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}

	if !strings.HasPrefix(port, ":") {
		port = fmt.Sprintf(":%s", port)
	}

	return e.Start(port)
}

const ctxHeader = "CTX_HEADER"

func SetHeader(ctx context.Context, header http.Header) context.Context {
	return context.WithValue(ctx, ctxHeader, header)
}

func GetHeader(ctx context.Context) (header http.Header, err error) {
	var has bool
	if header, has = ctx.Value(ctxHeader).(http.Header); !has {
		err = fmt.Errorf("not found header")
		return
	}
	return
}

func GetHeaderP(ctx context.Context) (header http.Header) {
	var err error
	if header, err = GetHeader(ctx); err != nil {
		panic(err)
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
