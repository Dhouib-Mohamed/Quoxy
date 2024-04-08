package router

import "github.com/gin-gonic/gin"

func Router() {
	router := gin.Default()
	healthRoutes(router)
	versionRoutes(router)
	subscriptionRoutes(router)
	tokenRoutes(router)
	notFoundRoutes(router)
	panic(router.Run(":8020"))
}
