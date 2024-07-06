package main

import (
	"github.com/gorilla/mux"
	"github.com/lbernardo/httperrors"
	"log"
	"net/http"
)

type customMessageErr struct {
	Message string `json:"message"`
}

func (c *customMessageErr) Error() string {
	return c.Message
}

func (c *customMessageErr) StatusCode() int {
	return http.StatusBadRequest
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping/{st}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		st := vars["st"]
		if st != "ok" {
			httperrors.GorillaHandlerError(w, &customMessageErr{Message: "custom error for ping"})
			return
		}
	})
	log.Fatal(http.ListenAndServe(":8081", r))
}
