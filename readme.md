# httperrors

This library has the functionality to handle Golang errors and convert them into HTTP status codes.

```shell
go get github.com/lbernardo/httperrors
```

## Usage

**gin**
```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lbernardo/httperrors"
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
	d := gin.Default()
	d.GET("/ping/:st", func(c *gin.Context) {
		st := c.Param("st")
		if st != "ok" {
			httperrors.GinHandleError(c, &customMessageErr{Message: "custom error for ping"})
			return
		}
		c.Status(http.StatusOK)
	})
	d.Run(":8081")
}
```

**gorilla/mux**
```go
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
```