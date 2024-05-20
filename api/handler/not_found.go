package handler

import (
	"api-authenticator-proxy/util/error_handler"
	routerError "api-authenticator-proxy/util/error_handler/router"
	"github.com/gin-gonic/gin"
)

func notFoundRoutes(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		error_handler.GinHandler(c, routerError.RouteNotFoundError())
	})
}
