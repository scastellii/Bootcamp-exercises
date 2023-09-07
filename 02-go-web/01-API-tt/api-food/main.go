package main

import (
	"encoding/json"
	"os"
)

type Product struct { //aca no hacen falta las etiquetas porque no voy a responder al usuario con esta estructura
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_Value"`
	IsPublished bool    `json:"is_Published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var products []Product
var LastId int = 500

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
