package router

import (
	"api-authenticator-proxy/src/database"
	"github.com/gin-gonic/gin"
)

func subscriptionRoutes(router *gin.Engine) {
	subscription := database.Subscription{}

	router.GET("/subscription/:id", func(c *gin.Context) {
		id := c.Param("id")
		subscription, err := subscription.GetById(id)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"result": subscription,
		})
	})

	router.GET("/subscription/name/:name", func(c *gin.Context) {
		name := c.Param("name")
		subscription, err := subscription.GetByName(name)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"result": subscription})
	})

	router.GET("/subscription", func(c *gin.Context) {
		subscriptions, err := subscription.GetAll()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if len(subscriptions) == 0 {
			c.JSON(200, gin.H{"result": []database.SubscriptionModel{}, "message": "No subscriptions found"})
			return
		}
		c.JSON(200, gin.H{"result": subscriptions})
	})

	router.POST("/subscription", func(c *gin.Context) {
		var newSubscription database.CreateSubscription
		if err := c.ShouldBindJSON(&newSubscription); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		result, err := subscription.Create(&newSubscription)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Subscription created", "result": result})
	})

	router.PUT("/subscription/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updateSubscription database.UpdateSubscription
		if err := c.ShouldBindJSON(&updateSubscription); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		result, err := subscription.Update(id, &updateSubscription)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Subscription updated", "result": result})
	})

	router.DELETE("/subscription/:id", func(c *gin.Context) {
		id := c.Param("id")
		result, err := subscription.Disable(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Subscription disabled", "result": result})
	})

	router.PATCH("/subscription/:id", func(c *gin.Context) {
		id := c.Param("id")
		result, err := subscription.Restore(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Subscription restored", "result": result})
	})
}
