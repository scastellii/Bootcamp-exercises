package main

import "fmt"

/*
Ejercicio 3 - A qué mes corresponde

Realizar una aplicación que reciba  una variable con el número del mes.
Según el número, imprimir el mes que corresponda en texto.
¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
Ej: 7, Julio.
Nota: Validar que el número del mes, sea correcto.
*/
func main() {
	var number = 7
	months := map[int]string{}
	months[1] = "Enero"
	months[2] = "Febrero"
	months[3] = "Marzo"
	months[4] = "Abril"
	months[5] = "Mayo"
	months[6] = "Junio"
	months[7] = "Julio"
	months[8] = "Agosto"
	months[9] = "Septiembre"
	months[10] = "Octubre"
	months[11] = "Noviembre"
	months[12] = "Diciembre"

	if number <= 0 || number > 12 {
		fmt.Println("El número de mes es incorrecto")
	} else {
		fmt.Println("El mes seleccionado es: ", months[number])
	}

}
