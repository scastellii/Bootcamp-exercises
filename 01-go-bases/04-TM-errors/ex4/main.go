package main

import (
	"errors"
	"fmt"
	my_errors "go-bases/01-go-bases/04-TM-errors/my-errors"
)

/*
Ejercicio 4 - Impuestos de salario #4
Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba
por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible
(el mensaje mostrado por consola deberá decir: “Error: el mínimo imponible es de 150.000 y el salario ingresado es de:
[salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
*/

func main() {
	//Ejercicio 4
	salary := 9000
	_, err := verifySalaryErrorsWithValue(salary)
	if errors.Is(err, my_errors.MyMinError2) {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}

func verifySalaryErrorsWithValue(salary int) (val bool, err error) {
	if salary <= 10000 {
		err = fmt.Errorf("%w y el salario ingresado es de: %v", my_errors.MyMinError2, salary)
		return
	} else {
		return true, nil
	}
}
