package medutils

import "errors"

var (

	// ErrInvalidObjId
	ErrInvalidObjId = errors.New("invalid object id")

	// ErrNotFound
	ErrNotFound = errors.New("mongo document not found")
)
