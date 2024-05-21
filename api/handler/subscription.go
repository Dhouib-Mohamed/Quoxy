package handler

import (
	"api-authenticator-proxy/internal/database"
	"api-authenticator-proxy/internal/models"
	"api-authenticator-proxy/util/error_handler"
	"github.com/gin-gonic/gin"
)

func subscriptionRoutes(router *gin.Engine) {
	subscription := database.Subscription{}

	router.GET("/subscription/:id",
		func(c *gin.Context) {
			id := c.Param("id")
			subscription, err := subscription.GetById(id)
			if err != nil {
				error_handler.GinHandler(c, err)
				return
			}
			c.JSON(200, gin.H{
				"result": subscription,
			})
		})

	router.GET("/subscription/name/:name",
		func(c *gin.Context) {
			name := c.Param("name")
			subscription, err := subscription.GetByName(name)
			if err != nil {
				error_handler.GinHandler(c, err)
				return
			}
			c.JSON(200, gin.H{"result": subscription})
		})

	router.GET("/subscription",
		func(c *gin.Context) {
			subscriptions, err := subscription.GetAll()
			if err != nil {
				error_handler.GinHandler(c, err)
				return
			}
			if len(subscriptions) == 0 {
				c.JSON(200, gin.H{"result": []models.SubscriptionModel{}, "message": "No subscriptions found"})
				return
			}
			c.JSON(200, gin.H{"result": subscriptions})
		})

	router.POST("/subscription",
		func(c *gin.Context) {
			var newSubscription models.CreateSubscription
			if err := c.BindJSON(&newSubscription); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			err := subscription.Create(&newSubscription)
			if err != nil {
				error_handler.GinHandler(c, err)
				return
			}
			c.JSON(201, gin.H{"message": "Subscription created"})
		})

	router.PUT("/subscription/:id",
		func(c *gin.Context) {
			id := c.Param("id")
			var updateSubscription models.UpdateSubscription
			if err := c.ShouldBindJSON(&updateSubscription); err != nil {
				c.JSON(400, gin.H{"error": "body fields not found"})
				return
			}
			err := subscription.Update(id, &updateSubscription)
			if err != nil {
				error_handler.GinHandler(c, err)
				return
			}
			c.JSON(200, gin.H{"message": "Subscription updated"})
		})

	router.DELETE("/subscription/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := subscription.Disable(id)
		if err != nil {
			error_handler.GinHandler(c, err)
			return
		}
		c.JSON(200, gin.H{"message": "Subscription disabled"})
	})

	router.PATCH("/subscription/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := subscription.Restore(id)
		if err != nil {
			error_handler.GinHandler(c, err)
			return
		}
		c.JSON(200, gin.H{"message": "Subscription restored"})
	})
}
