package httperrors

import (
	"github.com/gin-gonic/gin"
)

func GinHandleError(ctx *gin.Context, err error) {
	statusCode, message := Identify(err)
	ctx.JSON(statusCode, gin.H{"message": message})
}
