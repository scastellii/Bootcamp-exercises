package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}

func main() {
	p := Person{
		Name:     "",
		LastName: "",
	}
	router := gin.Default()

	router.POST("/saludo", func(c *gin.Context) {
		c.ShouldBind(&p)
		c.String(200, "Hola "+p.Name+" "+p.LastName)
	})
	router.Run()
}
