package error_handler

import (
	"api-authenticator-proxy/util/log"
	"fmt"
	"github.com/gin-gonic/gin"
)

type StatusError interface {
	GetError() (int, string)
}

func GinHandler(c *gin.Context, error StatusError) {
	code, err := error.GetError()
	c.JSON(code, gin.H{"error": err})
}

func CLIHandler(error StatusError) {
	code, err := error.GetError()
	log.Error(fmt.Errorf("error with code %d: %s", code, err))
}
