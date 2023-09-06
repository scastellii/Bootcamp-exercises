package main

import (
	"encoding/json"
	"os"
)

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_Value"`
	Is_published bool    `json:"is_Published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

var products []Product

func main() {
	path := "02-go-web/01-API-tt/products.json"
	products = getAllProducts(path)
	Init()

}

func getAllProducts(path string) []Product {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var products []Product
	err = json.Unmarshal(file, &products)
	if err != nil {
		return nil
	}
	return products
}

/*
Ejercicio 2 : Creando un servidor
Vamos a levantar un servidor en el puerto 8080. Para probar nuestros endpoints haremos uso de postman.
Crear una ruta /ping que debe respondernos con un string que contenga pong con el status 200 OK.
Crear una ruta /products que nos devuelva la lista de todos los productos en la slice.
Crear una ruta /products/:id que nos devuelva un producto por su id.
Crear una ruta /products/search que nos permita buscar por parámetro los productos cuyo precio sean mayor a un valor priceGt.
*/
