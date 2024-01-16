package wrGrpc

import (
	"fmt"
	"github.com/d3v-friends/go-tools/fn/fnPanic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"strings"
)

var ServerOptNoCredential = grpc.Creds(insecure.NewCredentials())

func Listen(sv *grpc.Server, port string) {
	var lis net.Listener

	if !strings.HasPrefix(port, ":") {
		port = fmt.Sprintf(":%s", port)
	}

	var err error
	if lis, err = net.Listen("tcp", port); err != nil {
		return
	}

	fmt.Printf("🚀 Server ready at http://localhost%s\n", port)
	fnPanic.On(sv.Serve(lis))
}
