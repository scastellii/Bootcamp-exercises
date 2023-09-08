package main

import (
	"go-bases/02-go-web/01-API-tt/api-food/cmd/server/handlers"
	"go-bases/02-go-web/01-API-tt/api-food/internal/products"
)

func main() {
	path := "02-go-web/01-API-tt/products.json"
	lastId := 500
	repo := products.NewRepositoryController(path, lastId)
	service := products.NewServiceController(repo)
	productController := handlers.NewProductController(service)
	Init(productController)
}
