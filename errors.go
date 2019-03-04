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
)
