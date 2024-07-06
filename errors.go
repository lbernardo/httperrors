package httperrors

import (
	"net/http"
)

type ErrorResponse interface {
	StatusCode() int
}

func Identify(err error) (int, string) {
	v, ok := err.(ErrorResponse)
	if !ok {
		return http.StatusInternalServerError, err.Error()
	}
	return v.StatusCode(), err.Error()
}
