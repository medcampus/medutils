package medutils

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"gopkg.in/mgo.v2/bson"
)

func ValidObjectId(objectId ...string) (error, bool) {
	for _, object := range objectId {
		if !bson.IsObjectIdHex(object) {
			return ConvertToGrpcErr(codes.InvalidArgument, fmt.Sprintf("Invalid ObjectId %v", object)), false
		}
	}
	return nil, true
}
