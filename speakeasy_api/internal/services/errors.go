package services

import "errors"

var (
	ErrNotImplemented = errors.New("method not implemented")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrNotFound       = errors.New("not found")
)
