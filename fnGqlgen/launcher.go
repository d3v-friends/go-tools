package fnGqlgen

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

func Run(port string, path string, handler *handler.Server) error {
	var e = echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTION"},
	}))
	e.POST(path, func(c echo.Context) error {
		var auth = c.Request().Header.Get("Authorization")
		var req = c.Request().WithContext(SetAuth(c.Request().Context(), auth))
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
