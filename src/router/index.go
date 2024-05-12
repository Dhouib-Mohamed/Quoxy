package router

import (
	"api-authenticator-proxy/src/utils/log"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	healthRoutes(router)
	versionRoutes(router)
	subscriptionRoutes(router)
	tokenRoutes(router)
	notFoundRoutes(router)
	log.Info("HTTP Server running on 8020")
	log.Fatal(router.Run(":8020"))
	return router
}
