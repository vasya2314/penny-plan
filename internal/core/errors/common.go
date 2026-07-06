package core_errors

import "errors"

var (
	ErrNotFound                = errors.New("not found")
	ErrBadRequest              = errors.New("bad request")
	ErrUnauthorized            = errors.New("unauthorized")
	ErrForbidden               = errors.New("forbidden")
	ErrInternalServerError     = errors.New("internal server error")
	ErrNotAcceptable           = errors.New("not acceptable")
	ErrRequestTimeout          = errors.New("request timeout")
	ErrConflict                = errors.New("conflict")
	ErrPreconditionFailed      = errors.New("precondition failed")
	ErrTooManyRequests         = errors.New("too many requests")
	ErrInternalServer          = errors.New("internal server error")
	ErrNotImplemented          = errors.New("not implemented")
	ErrBadGateway              = errors.New("bad gateway")
	ErrServiceUnavailable      = errors.New("service unavailable")
	ErrGatewayTimeout          = errors.New("gateway timeout")
	ErrHTTPVersionNotSupported = errors.New("http version not supported")
	ErrVariantAlsoNegotiates   = errors.New("also negotiates")
	ErrInsufficientStorage     = errors.New("insufficient storage")
	ErrLoopDetected            = errors.New("loop detected")
	ErrInvalidArgument         = errors.New("invalid argument")
)
