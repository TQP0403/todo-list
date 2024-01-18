package common

import (
	"net/http"
)

var CommonErr = map[int]string{
	http.StatusUnauthorized:        "Unauthorized",
	http.StatusBadRequest:          "Bad Request",
	http.StatusConflict:            "Conflict",
	http.StatusNotFound:            "Not Found",
	http.StatusTooManyRequests:     "Too Many Requests",
	http.StatusInternalServerError: "Server Error",
}

type AppError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
	LogError   string `json:"log"`
}

func (e *AppError) Error() string {
	return e.RootErr().Error()
}

func (e *AppError) RootErr() error {
	return e.Err
}

func NewAppError(statusCode int, message string, err error) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
		LogError:   err.Error(),
	}
}

func newCommonError(statusCode int, err error) *AppError {
	return NewAppError(statusCode, CommonErr[statusCode], err)
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
