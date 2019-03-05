package medutils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ConvertToGrpcErr(code codes.Code, message string) error {
	return status.Errorf(code, message)
}
