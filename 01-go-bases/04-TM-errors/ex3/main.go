package main

import (
	"errors"
	"fmt"
	my_errors "go-bases/01-go-bases/04-TM-errors/my-errors"
)

/*
Ejercicio 3 - Impuestos de salario #3

Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,
se implemente “errors.New()”.

*/

func main() {
	//Ejercicio 3
	salary := 9000
	_, err := verifySalaryErrors(salary)
	if errors.Is(err, my_errors.MyMinError) {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
func verifySalaryErrors(salary int) (val bool, err error) {
	if salary <= 10000 {
		err = fmt.Errorf("%w", my_errors.MyMinError)
		return
	} else {
		return true, nil
	}
}
