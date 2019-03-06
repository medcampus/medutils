package medutils

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func InitGrpc(port string, opts ...grpc.ServerOption) (*grpc.Server, net.Listener) {

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Errorf("Unable to create listener at port: %s, error: %v", port, err)
		return nil, nil
	}

	opts = append(opts, grpc.UnaryInterceptor(GRPC_Logging), grpc.UnaryInterceptor(GRPC_Recovery))

	srv := grpc.NewServer(opts...)

	reflection.Register(srv)

	return srv, listener
}
