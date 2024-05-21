package handler

import (
	"api-authenticator-proxy/util/error_handler"
	routerError "api-authenticator-proxy/util/error_handler/router"
	"api-authenticator-proxy/util/log"
	"fmt"
	"github.com/gin-gonic/gin"
)

func notFoundRoutes(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		log.Error(fmt.Errorf("the router %s is not found", c.FullPath()))
		error_handler.GinHandler(c, routerError.RouteNotFoundError(c.FullPath()))
	})
}
