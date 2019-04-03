package medutils

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func PutMetadata(userId, organisationId, userRole string) context.Context {
	header := metadata.New(map[string]string{"userid": userId, "organisationid": organisationId, "role": userRole})

	c := metadata.NewOutgoingContext(context.Background(), header)

	return c
}

func PullMetadata(ctx context.Context) (metadata.MD, error) {
	if md, ok := metadata.FromIncomingContext(ctx); !ok {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("unable to fetch metadata from context %v", ctx))
	} else {
		return md, nil
	}
}
