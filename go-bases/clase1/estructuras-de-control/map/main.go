package main

import "fmt"

/*
	Ejercicio 4 - Qué edad tiene...

Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
Según el siguiente mapa, ayudá a imprimir la edad de Benjamin.

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado también es necesario:
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.
*/
func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("La edad de benjamin es:", employees["Benjamin"])
	contador := 0
	for _, e := range employees {
		if e > 21 {
			contador++
		}
	}
	fmt.Println("La cantidad de empleados mayor que 21 años es: ", contador)
	employees["Federico"] = 25
	fmt.Println("Federico fue agregado con exito")
	delete(employees, "Pedro")
	fmt.Println("Pedro fue eliminado con exito")

}
