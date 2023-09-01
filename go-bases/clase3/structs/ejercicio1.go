package structs

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{}

func (p Product) Save(product Product) {
	Products = append(Products, product)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Println(product.Name)
	}
}

func (p Product) GetById(id int) Product {
	errorProduct := Product{}
	for _, product := range Products {
		if product.ID == id {
			return product
		}
	}
	return errorProduct
}
