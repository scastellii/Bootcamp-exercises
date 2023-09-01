package ejercicios

import "fmt"

/*
Ejercicio 3 - Calcular salario

Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría y que devuelva su salario.
*/
func EjecutarEjercicio3() {
	salario := calcularSalario(60, "A")
	fmt.Println("Cobra: ", salario)
	salario = calcularSalario(60, "B")
	fmt.Println("Cobra: ", salario)
	salario = calcularSalario(60, "C")
	fmt.Println("Cobra: ", salario)

}

func calcularSalario(minutos float64, categoria string) float64 {
	var salario float64 = 0

	switch categoria {
	case "C":
		salario = 1000 * (minutos / 60)
	case "B":
		salario = 1500 * (minutos / 60)
		salario += salario * 0.2

	case "A":
		salario = 3000 * (minutos / 60)
		salario += salario * 0.5
	default:
		return 0
	}
	return salario
}
