package medutils

import "errors"

var (
	// ErrGRPCPool indicates error while creating connection pool
	ErrGRPCPool = errors.New("error while creating connection pool")

	// ErrGRPCParse indicates error while parsing gRPC error object
	ErrGRPCParse = errors.New("not able to parse error returned")
)
