package medutils

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func InitGrpc(port string, serviceImpl interface{}, serviceInfo *grpc.ServiceDesc, opts ...grpc.ServerOption)  {

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Errorf("Unable to create listener at port: %s, error: %v", port, err)
		return
	}

	opts = append(opts, grpc.UnaryInterceptor(GRPC_Logging), grpc.UnaryInterceptor(GRPC_Recovery))

	srv := grpc.NewServer(opts...)

	srv.RegisterService(serviceInfo, serviceImpl)

	reflection.Register(srv)

	go func() {
		log.Infof("Starting FileService Server...\n Listening on port %v", port)
		if err := srv.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	shutdown(srv, listener)
}
