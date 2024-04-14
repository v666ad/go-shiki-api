package shikimori

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrTooManyRequests     = errors.New("too many requests")
	ErrInternalServer      = errors.New("internal server error")
	ErrUnauthorized        = errors.New("unauthorized access")
	ErrBadRequest          = errors.New("bad request")
	ErrForbidden           = errors.New("forbidden")
)
