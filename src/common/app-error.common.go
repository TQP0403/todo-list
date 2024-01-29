package common

import (
	"net/http"
)

var commonErr = map[int]string{
	http.StatusUnauthorized:        "Unauthorized",
	http.StatusBadRequest:          "Bad Request",
	http.StatusConflict:            "Conflict",
	http.StatusNotFound:            "Not Found",
	http.StatusTooManyRequests:     "Too Many Requests",
	http.StatusInternalServerError: "Server Error",
	http.StatusServiceUnavailable:  "Service Unavailable",
}

type AppError struct {
	StatusCode int    `json:"-"`
	Root       error  `json:"-"`
	Message    string `json:"message"`
	Err        string `json:"error"`
	Log        string `json:"log"`
}

// implement interface error
func (e *AppError) Error() string {
	return e.RootError().Error()
}

func (e *AppError) RootError() error {
	// parse AppError from interface error
	if val, ok := e.Root.(*AppError); ok {
		return val.RootError()
	}

	return e.Root
}

// error wrapper
func NewAppError(statusCode int, message string, err error) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Root:       err,
		Message:    "fail",
		Err:        message,
		Log:        err.Error(),
	}
}

func newCommonError(statusCode int, err error) *AppError {
	// do not wrap AppError
	if val, ok := err.(*AppError); ok {
		return val
	}

	return NewAppError(statusCode, commonErr[statusCode], err)
}

func NewUnauthorizedError(err error) *AppError {
	return newCommonError(http.StatusUnauthorized, err)
}

func NewBadRequestError(err error) *AppError {
	return newCommonError(http.StatusBadRequest, err)
}

func NewConflictError(err error) *AppError {
	return newCommonError(http.StatusConflict, err)
}

func NewNotFoundError(err error) *AppError {
	return newCommonError(http.StatusNotFound, err)
}

func NewTooManyRequestsError(err error) *AppError {
	return newCommonError(http.StatusTooManyRequests, err)
}

func NewInternalServerError(err error) *AppError {
	return newCommonError(http.StatusInternalServerError, err)
}

func NewServiceUnavailableError(err error) *AppError {
	return newCommonError(http.StatusServiceUnavailable, err)
}
