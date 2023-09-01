package ejercicios

import "testing"

func TestPromedio(t *testing.T) {
	t.Run("Calcula el promedio de la nota de los alumnos", func(t *testing.T) {
		//Arrange
		notas := []int{10, 5, 8, 7}
		//Act
		result := Promedio(notas...)
		//Assert
		if result != 30.0/4 {
			t.Fatal("Error al calcular el promedio")
		}
	})

}
