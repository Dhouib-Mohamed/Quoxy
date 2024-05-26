package handler

import (
	"api-authenticator-proxy/util/config"
	"api-authenticator-proxy/util/env"
	"api-authenticator-proxy/util/log"
	"github.com/gin-gonic/gin"
)

func Router() {
	if !config.GetIsRouterEnabled() {
		return
	}
	port := config.GetRouterPort()
	logLevel := env.GetLogLevel()
	if logLevel >= env.DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	healthRoutes(router)
	versionRoutes(router)
	subscriptionRoutes(router)
	tokenRoutes(router)
	notFoundRoutes(router)
	log.Info("HTTP Server running on port:", port)
	log.Fatal(router.Run(":" + port))
}
