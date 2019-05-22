package medutils

import "errors"

var (
	// ErrGRPCPool indicates error while creating connection pool
	ErrGRPCPool = errors.New("error while creating connection pool")

	// ErrGRPCParse indicates error while parsing gRPC error object
	ErrGRPCParse = errors.New("not able to parse error returned")

	// ErrEmptyAuthToken indicates empty auth cookie in request
	ErrEmptyAuthToken = errors.New("empty auth token")

	// ErrEmptyTokens indicates empty auth cookie in request
	ErrEmptyToken = errors.New("error empty token")

	// ErrEmptyGrpcHeaders indicates headers missing in grpc context
	ErrEmptyGrpcHeaders = errors.New("empty headers")
)
