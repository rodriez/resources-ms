package exceptions

import "net/http"

type ApiError struct {
	StatusCode int      `json:"status"`
	Message    string   `json:"message"`
	Errors     []string `json:"cause"`
}

func InternalError(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Error",
		Errors:     err,
	}
}

func NotFound(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusNotFound,
		Message:    "Bad Request",
		Errors:     err,
	}
}

func BadRequest(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusBadRequest,
		Message:    "Bad request",
		Errors:     err,
	}
}

func (e *ApiError) Error() string {
	return e.Message
}

func (e *ApiError) Code() int {
	return e.StatusCode
}

func (e *ApiError) Cause() []string {
	return e.Errors
}
