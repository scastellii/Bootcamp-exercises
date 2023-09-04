package my_errors

import (
	"errors"
	"fmt"
)

type MinNoEnoughError struct {
	Salary int
}

func (e MinNoEnoughError) Error() string {
	return fmt.Sprintf("Error: el salario ingresado (%d) no alcanza el mínimo imponible", e.Salary)
}

type HoursError struct {
	Hours float64
}

func (e HoursError) Error() string {
	return fmt.Sprintf("Error: el trabajador no puede haber trabajado menos de 80 horas mensuales y trabajo %v hs mensuales", e.Hours)
}

var MyMinError = errors.New("Error: el salario es menor a 10.000")
var MyMinError2 = errors.New("Error: el mínimo imponible es de 150.000")
