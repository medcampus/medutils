package medutils

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"reflect"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

var (
	js = &jsonpb.Marshaler{EnumsAsInts: true, EmitDefaults: true, OrigName: true}
)

// Logging interceptor.
func grpcLogging(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	if !viper.GetBool("app.logRequest") {
		log.Infof("calling %s, req=%s", info.FullMethod, marshal(req))
	} else {
		log.Infof("calling %s", info.FullMethod)
	}
	resp, err = handler(ctx, req)

	if err != nil {
		log.Errorf("method = %v, Error = %v", info.FullMethod, err)
	}

	log.Infof("finished %s, took=%v", info.FullMethod, time.Since(start))
	return resp, err
}

func marshal(x interface{}) string {
	if x == nil || reflect.ValueOf(x).IsNil() {
		return fmt.Sprintf("<nil>")
	}

	pb, ok := x.(proto.Message)
	if !ok {
		return fmt.Sprintf("Marshal to json error: not a proto message")
	}

	var buf bytes.Buffer
	if err := js.Marshal(&buf, pb); err != nil {
		return fmt.Sprintf("Marshal to json error: %s", err.Error())
	}
	return buf.String()
}
