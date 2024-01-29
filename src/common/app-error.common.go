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
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
}

// implement interface error
func (e *AppError) Error() string {
	return e.RootErr().Error()
}

func (e *AppError) RootErr() error {
	// parse AppError from interface error
	if val, ok := e.Err.(*AppError); ok {
		return val.RootErr()
	}

	return e.Err
}

func (err *AppError) GetErrorResponse() ErrorResponse {
	return ErrorResponse{Message: "error", Error: err.Message, Log: err.Error()}
}

// error wrapper
func NewAppError(statusCode int, message string, err error) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
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
