package main

import "go-bases/go-bases/clase3/structs"

func main() {
	person := structs.Person{
		ID:          123,
		Name:        "Santi",
		DateOfBirth: "30-04-1998",
	}

	employee := structs.Employee{
		Person:   person,
		Position: "ING",
		ID:       798,
	}

	employee.PrintEmployee()

	product := structs.Product{
		ID:          1,
		Name:        "GO",
		Price:       100.0,
		Description: "Lenguaje go",
		Category:    "Software",
	}
	product2 := structs.Product{
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
