package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("01-go-bases/04-TT-unknown/ex2/customers.txt")
	if err != nil {
		defer func() {
			fmt.Println("ejecución finalizada")
		}()
		panic(err)
	}
	fmt.Println("File open successfully")

	scanner := bufio.NewScanner(f)
	var text string
	for scanner.Scan() {
		text = scanner.Text()
	}
	text = ""
	if text == "" {
		defer func() {
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
		}()
		fmt.Println("ejecución finalizada")
		panic("Cannot read file")
	}

	defer func() {
		fmt.Println("ejecución finalizada")
	}()
}

/*
Ejercicio 2 - Imprimir datos

A continuación, vamos a crear un archivo “customers.txt” con información de los clientes del estudio.
Ahora que el archivo sí existe, el panic no debe ser lanzado.
Creamos el archivo “customers.txt” y le agregamos la información de los clientes.
Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos que contenga.
En el caso de no poder leerlo, se debe lanzar un “panic”.
Recordemos que siempre que termina la ejecución, independientemente del resultado, debemos tener un “defer” que nos indique que la ejecución finalizó.
También recordemos cerrar los archivos al finalizar su uso.
*/
