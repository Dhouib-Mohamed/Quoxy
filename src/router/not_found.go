package router

import "github.com/gin-gonic/gin"

func notFoundRoutes(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Route not found"})
	})
}
