package medutils

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
)

func ConvertToGrpcErr(code codes.Code, message string) error {
	return status.Errorf(code, message)
}

func GrpcLogError(context context.Context, err error) {
	logrus.Errorf("method = %v, Error = %v", context.Value("method"), err)
}

func ValidObjectId(objectId ...string) (error, bool) {
	for _, object := range objectId {
		if !bson.IsObjectIdHex(object) {
			return ConvertToGrpcErr(codes.InvalidArgument, fmt.Sprintf("Invalid ObjectId %v", object)), false
		}
	}
	return nil, true
}
