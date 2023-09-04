package main

import (
	"errors"
	"fmt"
	my_errors "go-bases/01-go-bases/04-TM-errors/my-errors"
)

/*
Ejercicio 2 - Impuestos de salario #2
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje
“Error: el salario es menor a 10.000" y lanzalo en caso de que “salary” sea menor o igual a  10.000.
La validación debe ser hecha con la función “Is()” dentro del “main”.
*/
func main() {

	//Ejercicio 2
	salary := 9000
	_, err := verifySalary(salary)
	if errors.Is(err, my_errors.MinNoEnoughError{Salary: salary}) {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}

func verifySalary(salary int) (val bool, err error) {
	if salary <= 10000 {
		err = my_errors.MinNoEnoughError{Salary: salary}
		return
	} else {
		return true, nil
	}
}
