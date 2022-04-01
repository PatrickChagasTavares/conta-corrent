package model

import (
	"fmt"
	"net/http"
)

type Error struct {
	HTTPCode int         `json:"-"`
	Message  string      `json:"message"`
	Detail   interface{} `json:"detail,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %v - message: %v - detail: %v", e.HTTPCode, e.Message, e.Detail)
}

// NewError cria um novo erro
func NewError(httpCode int, message string, detail interface{}) error {
	return &Error{
		HTTPCode: httpCode,
		Message:  message,
		Detail:   detail,
	}
}

// GetHTTPCode retorna o código http do erro
func GetHTTPCode(err error) int {
	e, ok := err.(*Error)
	if !ok {
		return http.StatusInternalServerError
	}
	return e.HTTPCode
}
