package main

import (
	structs2 "go-bases/01-go-bases/03-clase3/structs"
)

func main() {
	person := structs2.Person{
		ID:          123,
		Name:        "Santi",
		DateOfBirth: "30-04-1998",
	}

	employee := structs2.Employee{
		Person:   person,
		Position: "ING",
		ID:       798,
	}

	employee.PrintEmployee()

	product := structs2.Product{
		ID:          1,
		Name:        "GO",
		Price:       100.0,
		Description: "Lenguaje go",
		Category:    "Software",
	}
	product2 := structs2.Product{
		ID:          2,
		Name:        "Intellij",
		Price:       100.0,
		Description: "IDE",
		Category:    "Software",
	}
	product.Save(product2)
	product.Save(product)
	product.GetById(1)
	product.GetAll()

}
