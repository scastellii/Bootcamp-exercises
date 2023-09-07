package main

import "go-bases/02-go-web/01-API-tt/api-food/internal/product"

func main() {
	path := "02-go-web/01-API-tt/products.json"
	lastId := 500
	products := product.GetAllProducts(path)
	Init(product.NewControllerProducts(products, lastId))
}
