package apperror

import "net/http"

func BadRequest(message string) AppError {
	return New(http.StatusBadRequest, message)
}

func IsValid(err error) (AppError, bool) {
	appErr, ok := err.(AppError)
	return appErr, ok
}
