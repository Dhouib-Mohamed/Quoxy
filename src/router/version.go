package router

import "github.com/gin-gonic/gin"

var version = "0.0.0"

func versionRoutes(router *gin.Engine) {
	router.GET("/version", func(c *gin.Context) {
		c.JSON(200, gin.H{"version": version})
	})
}
