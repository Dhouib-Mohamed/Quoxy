package router

import (
	"api-authenticator-proxy/src/utils/error_handler"
	routerError "api-authenticator-proxy/src/utils/error_handler/router"
	"github.com/gin-gonic/gin"
)

func notFoundRoutes(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		error_handler.GinHandler(c, routerError.RouteNotFoundError())
	})
}
