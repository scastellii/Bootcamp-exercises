package main

import (
	"github.com/gin-gonic/gin"
	"go-bases/02-go-web/01-API-tt/api-food/internal/product"
)

func Init(c *product.ControllerProducts) {
	router := gin.Default()
	router.GET("/ping", ping)
	pr := router.Group("/products")
	pr.GET("", c.GetProducts())
	pr.GET("/:id", c.GetProduct())
	pr.POST("", c.SaveProduct())
	router.Run()
}

func ping(c *gin.Context) {
	c.String(200, "pong")
}
