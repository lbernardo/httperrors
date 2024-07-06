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
