package medutils

import "errors"

var (

	// ErrInvalidObjId indicates invalid mongo object id
	ErrInvalidObjId = errors.New("invalid object id")

	// ErrNotFound indicates document not found
	ErrNotFound = errors.New("mongo document not found")

	// ErrReadClaims indicates error while reading claims fro token
	ErrReadClaims = errors.New("error reading jwt claims")

	// ErrTokenInvalid indicates
	ErrTokenInvalid = errors.New("invalid JWT token")

	// ErrAuthTokenEmpty indicates auth token empty in cookie
	ErrAuthTokenEmpty = errors.New("error auth token empty in cookie")

	// ErrInvalidAuthHeader indicates user_manager header is invalid, could for example have the wrong Realm name
	ErrInvalidAuthHeader = errors.New("authorization header is invalid")

	//ErrInvalidArguments indicates arguments are not appropriate
	ErrInvalidArgs = errors.New("invalid arguments")

	//ErrEmailValidaton indicates input email not valid
	ErrEmailValidation = errors.New("failed email validation")

	// ErrGRPCPool indicates error while creating connection pool
	ErrGRPCPool = errors.New("error while creating connection pool")

	ErrGRPCParse = errors.New("not able to parse error returned")
)
