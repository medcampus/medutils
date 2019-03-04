package medutils

import (
	"bytes"
	"fmt"
	"reflect"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	js = &jsonpb.Marshaler{EnumsAsInts: true, EmitDefaults: true, OrigName: true}
)

// Logging interceptor.
func GRPC_Logging(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	log.Infof("calling %s, req=%s", info.FullMethod, marshal(req))
	resp, err = handler(ctx, req)
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
