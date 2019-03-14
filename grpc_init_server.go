package medutils

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strings"
)

func InitGrpc(opts ...grpc.ServerOption) (*grpc.Server, net.Listener) {

	InitConfig()

	InitLog(viper.GetString("app.logPath"))

	port := viper.GetString("app.port")
	if !strings.Contains(viper.GetString("app.port"), ":") {
		port = ":" + port
	}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Errorf("Unable to create listener at port: %s, error: %v", port, err)
		return nil, nil
	}

	opts = append(opts, grpc.UnaryInterceptor(unaryInterceptorChain(grpcLogging, grpcRecovery)))

	srv := grpc.NewServer(opts...)

	reflection.Register(srv)

	return srv, listener
}
