package controllers

import (
	"go-store/cmd/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {

	router.GET("/up", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "happy"})
	})

	auth := router.Group("/auth")
	{
		auth.POST("/signup", Signup)
		auth.POST("/login", Login)
	}

	products := router.Group("/products")
	{
		products.GET("", ListProduct)
		products.GET("/:id", GetProduct)

	}
	rules := router.Group("/rules")
	{
		rules.GET("", ListProductCartRules)
	}

	orders := router.Group("/orders")
	orders.POST("", CreateOrder)
	orders.Use(middlewares.EnforceAuthenticatedMiddleware())
	{
		orders.GET("", ListOrders)
		orders.GET("/:id", GetOrder)

	}

}
