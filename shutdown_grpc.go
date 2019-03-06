package medutils

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

func waitForInterrupt() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c
}

func shutdown(s *grpc.Server, l net.Listener) {
	// wait for interupt signal
	waitForInterrupt()

	logrus.Infof("Stopping gRPC server: %v", s.GetServiceInfo())
	s.GracefulStop()

	logrus.Infof("Closing listener at: %v", l.Addr())
	l.Close()
}
