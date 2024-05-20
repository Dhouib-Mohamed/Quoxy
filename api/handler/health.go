package handler

import "github.com/gin-gonic/gin"

func healthRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
