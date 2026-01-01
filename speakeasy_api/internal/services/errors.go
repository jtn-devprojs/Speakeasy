package services

import "errors"

var (
	// ErrNotImplemented indicates a method is not yet implemented
	ErrNotImplemented = errors.New("method not implemented")
	// ErrUnauthorized indicates an authentication failure
	ErrUnauthorized = errors.New("unauthorized")
	// ErrNotFound indicates a resource was not found
	ErrNotFound = errors.New("not found")
)
