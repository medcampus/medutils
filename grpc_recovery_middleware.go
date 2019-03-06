package medutils

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"runtime"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
)

const (
	MAXSTACKSIZE = 4096
)

// Recovery interceptor to handle grpc panic.
func grpc_recovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// recovery func
	defer func() {
		if r := recover(); r != nil {
			// log stack
			stack := make([]byte, MAXSTACKSIZE)
			stack = stack[:runtime.Stack(stack, false)]
			log.Infof("panic grpc invoke: %s, err=%v, stack:\n%s", info.FullMethod, r, string(stack))

			err = status.Errorf(codes.Internal, "panic error: %v", r)
		}
	}()

	return handler(ctx, req)
}
