package router

import (
	"api-authenticator-proxy/src/database"
	"api-authenticator-proxy/src/database/models"
	"api-authenticator-proxy/src/utils/error_handler"
	"github.com/gin-gonic/gin"
)

func tokenRoutes(router *gin.Engine) {

	token := database.Token{}
	tokenRouter := router.Group("/token")

	tokenRouter.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		token, err := token.GetById(id)
		if err != nil {
			error_handler.GinHandler(c, err)
			return
		}
		c.JSON(200, gin.H{"result": token})
	})

	tokenRouter.GET("/", func(c *gin.Context) {
		tokens, err := token.GetAll()
		if err != nil {
			error_handler.GinHandler(c, err)
			return
		}
		if len(tokens) == 0 {
			c.JSON(200, gin.H{"result": []models.TokenModel{}, "message": "No tokens found"})
			return
		}
		c.JSON(200, gin.H{"result": tokens})
	})

	tokenRouter.POST("/", func(c *gin.Context) {
		var newToken models.CreateToken
		if err := c.ShouldBindJSON(&newToken); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		res, err := token.Create(&newToken)
		if err != nil {
			error_handler.GinHandler(c, err)
			return
		}
		c.JSON(201, gin.H{"message": "Token created", "result": res})
	})

	tokenRouter.PUT("/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updateToken models.UpdateToken
		if err := c.ShouldBindJSON(&updateToken); err != nil {
			c.JSON(400, gin.H{"error": ""})
			return
		}
		err := token.Update(id, &updateToken)
		if err != nil {
			error_handler.GinHandler(c, err)
			return
		}
		c.JSON(200, gin.H{"message": "Token updated"})
	})

	tokenRouter.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := token.Disable(id)
		if err != nil {
			error_handler.GinHandler(c, err)
			return
		}
		c.JSON(200, gin.H{"message": "Token deleted"})
	})
}
