package fnGrpc

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"strings"
)

func NewServer() *grpc.Server {
	panic("not implement")
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
