package httperrors

import (
	"net/http"
)

func GorillaHandlerError(w http.ResponseWriter, err error) {
	statusCode, message := Identify(err)
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}
