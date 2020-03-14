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
	products.Use(middlewares.EnforceAuthenticatedMiddleware())
	{
		products.GET("", ListProduct)
		products.GET("/:id", GetProduct)
	}

	orders := router.Group("/orders")
	orders.Use(middlewares.EnforceAuthenticatedMiddleware())
	{
		orders.GET("", ListOrders)
		orders.POST("", CreateOrder)
		orders.GET("/:id", GetOrder)

	}

}
