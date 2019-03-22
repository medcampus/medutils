package medutils

import "errors"

var (
	// ErrGRPCPool indicates error while creating connection pool
	ErrGRPCPool = errors.New("error while creating connection pool")

	// ErrGRPCParse indicates error while parsing gRPC error object
	ErrGRPCParse = errors.New("not able to parse error returned")

	// ErrEmptyAuthToken indicates empty auth cookie in request
	ErrEmptyAuthToken = errors.New("empty auth token")
)
