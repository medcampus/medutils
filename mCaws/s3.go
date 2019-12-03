package mCaws

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type medcampusS3 interface {
	SignedUrl(context.Context, string, string) (string, error)
}

type S3Object struct {
	Session *s3.S3
}

func (s3Obj *S3Object) SignedUrl(ctx context.Context, bucket string, key string) (string, error) {
	req, _ := s3Obj.Session.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	signedUrl, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", status.Errorf(codes.Internal, "Failed to sign request, error %v", err)
	}

	return signedUrl, nil
}
