package ejercicios

import "fmt"

const (
	Minimum = "minimum"
	Average = "average"
	Maximum = "maximum"
)

func EjecutarEjercicio4(notas ...int) {
	fx, msg := realizarOperacion(Maximum)
	if msg != "Error" {
		resultado := fx(notas...)
		fmt.Println("El resultado de aplicar", Maximum, " es :", resultado)
	}
	fx, msg = realizarOperacion(Minimum)
	if msg != "Error" {
		resultado := fx(notas...)
		fmt.Println("El resultado de aplicar", Minimum, " es :", resultado)

	}
	fx, msg = realizarOperacion(Average)
	if msg != "Error" {
		resultado := fx(notas...)
		fmt.Println("El resultado de aplicar", Average, " es :", resultado)

	}
	fx, msg = realizarOperacion("HOLA")
	if msg != "Error" {
		resultado := fx(notas...)
		fmt.Println("El resultado de aplicar HOLA es :", resultado)

	}

}

func realizarOperacion(operation string) (aplicar func(values ...int) float64, mensaje string) {
	switch operation {
	case Maximum:
		aplicar = minValue
	case Average:
		aplicar = averageValue
	case Minimum:
		aplicar = minValue
	default:
		mensaje = "Error"
	}

	return

}

func minValue(values ...int) float64 {
	result := 0.0
	for i, value := range values {
		castValue := float64(value)
		if i == 0 {
			result = float64(castValue)
		} else {
			if castValue < result {
				result = castValue
			}
		}
	}
	return result
}

func maxValue(values ...int) float64 {
	result := 0.0
	for i, value := range values {
		castValue := float64(value)
		if i == 0 {
			result = float64(castValue)
		} else {
			if castValue > result {
				result = castValue
			}
		}
	}
	return result
}

func averageValue(values ...int) float64 {
	return Promedio(values...)
}

/*
Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes
de un curso.
Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros
y devuelva el cálculo que se indicó en la función anterior.
Ejemplo:


...
minFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)
...
minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
*/
