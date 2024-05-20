package error_handler

import "github.com/gin-gonic/gin"

type StatusError interface {
	GetError() (int, string)
}

func GinHandler(c *gin.Context, error StatusError) {
	code, err := error.GetError()
	c.JSON(code, gin.H{"error": err})
}
