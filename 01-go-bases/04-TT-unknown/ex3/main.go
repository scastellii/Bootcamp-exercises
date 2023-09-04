package main

import (
	"errors"
	"fmt"
)

type cliente struct {
	legajo    string
	nombre    string
	dni       string
	telefono  string
	domicilio string
}

func main() {
	var clients []cliente
	clients = addClients(clients)
	newClient := cliente{
		domicilio: "Aasdasd",
		dni:       "123",
		nombre:    "Pedasdasdro",
		telefono:  "21121313",
		legajo:    "",
	}

	existFlag := false
	for _, client := range clients {
		if client.dni == newClient.dni {
			defer func() {
				recover()
				existFlag = true
				fmt.Println("el cliente ya existe")
				fmt.Println("se detectaron varios errores en tiempo de ejecución")
				fmt.Println("Fin de la ejecución")
			}()
			panic("El cliente ya existe")
		}
	}
	if !existFlag {
		clients = append(clients, newClient)
		if result, err := validateData(newClient); !result {
			if err != nil {
				defer func() {
					recover()
					fmt.Println(err)
					fmt.Println("se detectaron varios errores en tiempo de ejecución")
					fmt.Println("Fin de la ejecución")
				}()
				panic(err)
			}
		} else {
			fmt.Println("El cliente fue agregado correctamente")
		}
	}
}

func validateData(client cliente) (bool, error) {
	result := client.telefono != "" && client.legajo != "" && client.nombre != "" && client.domicilio != "" && client.dni != ""
	if !result {
		return false, errors.New("los datos ingresados no son validos")
	}
	return result, nil
}

func addClients(clients []cliente) []cliente {
	client := cliente{
		domicilio: "A",
		dni:       "123",
		nombre:    "Pedro",
		telefono:  "21113",
		legajo:    "12343423",
	}
	clients = append(clients, client)
	return clients
}

/*
Ejercicio 3 - Registrando clientes
El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes. Los datos requeridos son:
Legajo
Nombre
DNI
Número de teléfono
Domicilio


Tarea 1: Antes de registrar a un cliente, debés verificar si el mismo ya existe. Para ello, necesitás leer los datos de un array.
En caso de que esté repetido, debes manipular adecuadamente el error como hemos visto hasta aquí. Ese error deberá:
1.- generar un panic;
2.- lanzar por consola el mensaje: “Error: el cliente ya existe”, y continuar con la ejecución del programa normalmente.


Tarea 2: Luego de intentar verificar si el cliente a registrar ya existe, desarrollá una función para validar que todos los datos a registrar de un cliente
contienen un valor distinto de cero.
Esta función debe retornar, al menos, dos valores. Uno de ellos tendrá que ser del tipo error para el caso de que se ingrese por parámetro algún valor cero
(recordá los valores cero de cada tipo de dato, ej: 0, “”, nil).
Tarea 3: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes:
“Fin de la ejecución” y “Se detectaron varios errores en tiempo de ejecución”. Utilizá defer para cumplir con este requerimiento.

Requerimientos generales:
Utilizá recover para recuperar el valor de los panics que puedan surgir
Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error.
Generá algún error, personalizandolo a tu gusto utilizando alguna de las funciones de Go (realiza también la validación pertinente para el caso de error retornado).

*/
