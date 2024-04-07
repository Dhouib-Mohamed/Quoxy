package router

import (
	"api-authenticator-proxy/src/database"
	"github.com/gin-gonic/gin"
)

func tokenRoutes(router *gin.Engine) {

	token := database.Token{}
	tokenRouter := router.Group("/token")

	tokenRouter.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		token, err := token.GetById(id)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"result": token})
	})

	tokenRouter.GET("/", func(c *gin.Context) {
		tokens, err := token.GetAll()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if len(tokens) == 0 {
			c.JSON(200, gin.H{"result": []database.TokenModel{}, "message": "No tokens found"})
			return
		}
		c.JSON(200, gin.H{"result": tokens})
	})

	tokenRouter.POST("/", func(c *gin.Context) {
		var newToken database.CreateToken
		if err := c.ShouldBindJSON(&newToken); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		result, err := token.Create(&newToken)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Token created", "result": result})
	})

	tokenRouter.PUT("/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updateToken database.UpdateToken
		if err := c.ShouldBindJSON(&updateToken); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		result, err := token.Update(id, &updateToken)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Token updated", "result": result})
	})

	tokenRouter.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		result, err := token.Disable(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Token deleted", "result": result})
	})
}
