package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Init() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProduct)

	router.Run()

}

func ping(c *gin.Context) {
	c.String(200, "pong")

}

func getProducts(c *gin.Context) {
	if qParam, ok := c.GetQuery("priceGt"); ok {
		priceToSearch, _ := strconv.Atoi(qParam)
		var filteredProducts []Product
		for _, product := range products {
			if product.Price > float64(priceToSearch) {
				filteredProducts = append(filteredProducts, product)
			}
		}
		c.JSON(http.StatusOK, &filteredProducts)
		return
	}
	c.JSON(http.StatusOK, products)
}

func getProduct(c *gin.Context) {
	idToSearch := c.Param("id")
	idParse, _ := strconv.Atoi(idToSearch)
	for _, product := range products {
		if product.Id == idParse {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(500, "No se encontró el producto")
	return
}

/*
Ejercicio 2 : Creando un servidor
Vamos a levantar un servidor en el puerto 8080. Para probar nuestros endpoints haremos uso de postman.
Crear una ruta /ping que debe respondernos con un string que contenga pong con el status 200 OK.
Crear una ruta /products que nos devuelva la lista de todos los productos en la slice.
Crear una ruta /products/:id que nos devuelva un producto por su id.
Crear una ruta /products/search que nos permita buscar por parámetro los productos cuyo precio sean mayor a un valor priceGt.
*/
