package wrGrpc

import (
	"crypto/tls"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"strings"
)

func NewServer(opts ...grpc.ServerOption) *grpc.Server {
	opts = append(opts, grpc.Creds(credentials.NewTLS(&tls.Config{
		ClientAuth: tls.NoClientCert,
	})))

	return grpc.NewServer(
		opts...,
	)
}

func Listen(sv *grpc.Server, port string) (err error) {
	var lis net.Listener

	if !strings.HasPrefix(port, ":") {
		port = fmt.Sprintf(":%s", port)
	}

	if lis, err = net.Listen("tcp", port); err != nil {
		return
	}

	fmt.Printf("🚀 Server ready at http://localhost%s\n", port)
	return sv.Serve(lis)
}
