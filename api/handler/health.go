package handler

import (
	"api-authenticator-proxy/util/log"
	"github.com/gin-gonic/gin"
)

func healthRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		log.Debug("Checking Http Server Status ...")
		c.JSON(200, gin.H{"status": "ok"})
	})
}
