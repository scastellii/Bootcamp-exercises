package ejercicios

import "fmt"

/*
Ejercicio 2 - Calcular promedio
Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. No se pueden introducir notas negativas.
*/
func EjecutarEjercicio2() {
	prom := Promedio(10, 5, 8, 7)
	fmt.Println("El promedio de las notas es de: ", prom)
	prom = Promedio(9, 9, 8, 4)
	fmt.Println("El promedio de las notas es de: ", prom)
	prom = Promedio(7, 8, 9, 10, 6, 10, 10, 10, 10)
	fmt.Println("El promedio de las notas es de: ", prom)
}

func Promedio(notas ...int) float64 {
	var sum float64 = 0
	for _, nota := range notas {
		sum += float64(nota)
	}
	return sum / float64(len(notas))
}
