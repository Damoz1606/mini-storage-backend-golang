package apperror

import (
	"encoding/json"
	"fmt"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const errFmt = "%d - %s"

func (a AppError) Error() string {
	return fmt.Sprintf(errFmt, a.Code, a.Message)
}

func (a AppError) Marshal() []byte {
	e, _ := json.Marshal(a)
	return e
}

func New(code int, message string) AppError {
	return AppError{
		Code:    code,
		Message: message,
	}
}
