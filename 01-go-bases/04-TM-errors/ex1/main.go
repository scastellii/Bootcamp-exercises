package main

import (
	"fmt"
	my_errors "go-bases/01-go-bases/04-TM-errors/my-errors"
)

/*
Ejercicio 1 - Impuestos de salario #1
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje
“Error: el salario ingresado no alcanza el mínimo imponible" y lanzalo en caso de que “salary” sea menor a 150.000.
De lo contrario, tendrás que imprimir por consola el mensaje “Debe pagar impuesto”.
*/

func main() {
	//Ejercicio 1
	salary := 140000
	if salary < 150000 {
		enoughError := my_errors.MinNoEnoughError{Salary: salary}
		err := enoughError.Error()
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
