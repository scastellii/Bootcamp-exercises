package ejercicios

import "fmt"

/*
Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará además un 10 % (27% en total).
*/
func EjecutarEjercicio1() {
	var sueldo float64 = 60000
	descuento := calcularImpuesto(sueldo)
	fmt.Println("El impuesto de un sueldo de ", sueldo, "es igual a:", descuento)
	sueldo = 160000
	descuento = calcularImpuesto(sueldo)
	fmt.Println("El impuesto de un sueldo de ", sueldo, "es igual a:", descuento)
	sueldo = 10000
	descuento = calcularImpuesto(10000)
	fmt.Println("El impuesto de un sueldo de ", sueldo, "es igual a:", descuento)
}

func calcularImpuesto(salario float64) float64 {
	var descuento float64 = 0
	if salario > 50000 {
		descuento = salario * 0.17
	}
	if salario > 150000 {
		descuento += salario * 0.1
	}

	return descuento

}
