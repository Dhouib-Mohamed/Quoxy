package handler

import (
	"github.com/gin-gonic/gin"
	"os"
)

func versionRoutes(router *gin.Engine) {
	router.GET("/version", func(c *gin.Context) {
		version, _ := os.ReadFile("version.txt")
		c.JSON(200, gin.H{"version": string(version)})
	})
}
