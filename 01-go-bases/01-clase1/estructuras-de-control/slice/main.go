package main

import "fmt"

func main() {
	var (
		age                = 25
		antiguedad         = 2
		empleado           = true
		sueldo     float64 = 150000
	)

	switch {
	case age <= 22:
		fmt.Println("Es menor de 23 años")
	case !empleado:
		fmt.Println("El usuario no está trabajando")
	case antiguedad < 1:
		fmt.Println("El usuario no tiene suficiente antiguedad")
	case sueldo < 100000:
		fmt.Print("El sueldo no es suficiente")
	default:
		fmt.Println("Puede recibir presetamo")
	}

}
