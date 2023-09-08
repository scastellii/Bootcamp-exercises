package main

import (
	"github.com/gin-gonic/gin"
	"go-bases/02-go-web/01-API-tt/api-food/cmd/server/handlers"
)

func Init(c *handlers.ProductController) {
	router := gin.Default()
	router.GET("/ping", ping)
	pr := router.Group("/products")
	pr.GET("", c.GetAll())
	pr.GET("/:id", c.GetById())
	pr.POST("", c.Store())
	pr.PUT("/:id", c.Update())
	pr.DELETE("/:id", c.Delete())
	router.Run()
}

func ping(c *gin.Context) {
	c.String(200, "pong")
}
