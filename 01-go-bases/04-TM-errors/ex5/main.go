package main

import (
	"errors"
	"fmt"
	my_errors "go-bases/01-go-bases/04-TM-errors/my-errors"
)

/*
Ejercicio 5 -  Impuestos de salario #5
Vamos a hacer que nuestro programa sea un poco más complejo y útil.
Desarrollá las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).
En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10 % en concepto de impuesto.
En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo,
la función debe devolver un error. El mismo tendrá que indicar
“Error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
*/

func main() {
	//Ejercicio 5
	hours := 10.0
	valueHour := 1000.0
	salary, err := calculateSalary(hours, valueHour)
	if errors.Is(err, my_errors.HoursError{hours}) {
		fmt.Println(err)
	} else {
		fmt.Println("El trabajador cobra $", salary)
	}
}

func calculateSalary(workHour float64, valueHour float64) (salary float64, err error) {
	finalSalary := workHour * valueHour
	if workHour < 80.0 || workHour < 0.0 {
		err = my_errors.HoursError{Hours: workHour}
		return
	} else {
		if salary >= 150000 {
			return finalSalary - finalSalary*0.1, nil
		} else {
			return finalSalary, err
		}
	}
}
