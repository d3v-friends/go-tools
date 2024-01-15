package fnGrpc

import (
	"context"
	"github.com/d3v-friends/go-tools/fnLogger"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"time"
)

type FnSetContext func(ctx context.Context) context.Context

func Interceptor(fn FnSetContext, logger fnLogger.IfLogger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		var requestAt = time.Now()
		var requestLogger = logger.WithFields(fnLogger.Fields{
			"requestId": uuid.NewString(),
			"method":    info.FullMethod,
			"requestAt": time.Now(),
		})
		requestLogger.Trace("requested")

		ctx = fn(ctx)

		defer func() {
			var responseAt = time.Now()
			var responseLogger = requestLogger.
				WithFields(fnLogger.Fields{
					"responseAt": responseAt,
					"durations":  responseAt.UnixMilli() - requestAt.UnixMilli(),
				})

			if err == nil {
				responseLogger.
					Trace("responded")
			} else {
				responseLogger.
					Error("error: err=%s", err.Error())
			}
		}()

		return handler(ctx, req)
	}
}
