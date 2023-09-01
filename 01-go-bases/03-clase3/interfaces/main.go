package main

import (
	"errors"
	"fmt"
	"go-bases/01-go-bases/03-clase3/interfaces/products"
	"go-bases/01-go-bases/03-clase3/interfaces/students-registry"
)

const (
	Little = "Little"
	Medium = "Medium"
	Big    = "Big"
)

func main() {
	// Ejercicio estudiantes
	student := students.Alumnos{
		Nombre:   "pepe",
		Apellido: "grillo",
		DNI:      "1234",
		Fecha:    "23/07/2023",
	}
	student.Detalle()

	//Ejercicio productos
	pr := factory(Little, 300.0)
	fmt.Println("El precio total del producto", Little, "es:", pr.Price())
	pr = factory(Medium, 300.0)
	fmt.Println("El precio total del producto", Medium, "es:", pr.Price())
	pr = factory(Big, 300.0)
	fmt.Println("El precio total del producto", Big, "es:", pr.Price())
	pr = factory("HOLA", 300.0)
}

func factory(typeProduct string, precio float64) products.Product {
	switch typeProduct {
	case Little:
		return products.LittleProduct{Cost: precio}
	case Medium:
		return products.MediumProduct{Cost: precio, MaintenanceCost: 0.03}
	case Big:
		return products.BigProduct{Cost: precio, MaintenanceCost: 0.06, AdditionalCost: 2500.0}
	default:
		errors.New("No existe el type: " + typeProduct)
	}
	return nil
}
