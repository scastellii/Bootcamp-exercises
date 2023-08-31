package ejercicios

import "fmt"

const (
	Dog     = "dog"
	Cat     = "cat"
	Spider  = "spider"
	Hamster = "hamster"
)

func EjecutarEjercicio5(totalAnimales int) {
	fx, msg := calcularAlimento(Dog)
	if msg != "Error" {
		resultado := fx(totalAnimales)
		fmt.Println("Se necesitan", resultado, "gramos de alimento para", Dog)
	}
	fx, msg = calcularAlimento(Cat)
	if msg != "Error" {
		resultado := fx(totalAnimales)
		fmt.Println("Se necesitan", resultado, "gramos de alimento para", Cat)

	}
	fx, msg = calcularAlimento(Spider)
	if msg != "Error" {
		resultado := fx(totalAnimales)
		fmt.Println("Se necesitan", resultado, "gramos de alimento para", Spider)

	}
	fx, msg = calcularAlimento(Hamster)
	if msg != "Error" {
		resultado := fx(totalAnimales)
		fmt.Println("Se necesitan", resultado, "gramos de alimento para", Hamster)

	}
	fx, msg = calcularAlimento("HOLA")
	if msg != "Error" {
		resultado := fx(totalAnimales)
		fmt.Println("El resultado de aplicar HOLA es :", resultado)

	}

}

func calcularAlimento(animal string) (apply func(int) int, message string) {
	switch animal {
	case Spider:
		apply = alimentoSpider
	case Dog:
		apply = alimentoDog
	case Cat:
		apply = alimentoCat
	case Hamster:
		apply = alimentoHamster
	default:
		message = "Error"
	}
	return
}

func alimentoDog(cantidad int) int {
	return cantidad * 10000
}

func alimentoCat(cantidad int) int {
	return cantidad * 5000
}

func alimentoSpider(cantidad int) int {
	return cantidad * 150
}

func alimentoHamster(cantidad int) int {
	return cantidad * 250
}

/*
Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.
Tienen los siguientes datos:
Perro 10 kg de alimento.
Gato 5 kg de alimento.
Hamster 250 g de alimento.
Tarántula 150 g de alimento.
Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne
una función y un mensaje (en caso que no exista el animal).
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
Ejemplo:


const (
	dog    = "dog"
	cat    = "cat"
)


...


animalDog, msg := animal(dog)
animalCat, msg := animal(cat)


...


var amount float64
amount += animalDog(10)
amount += animalCat(10)

*/
