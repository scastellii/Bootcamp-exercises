package ejercicios

import "testing"

func TestCalcularSalario(t *testing.T) {
	t.Run("Calcula el salario para categoria A", func(t *testing.T) {
		//Arrange
		var minutos = 60.0
		//Act
		result := calcularSalario(minutos, "A")
		//Assert
		if result != 4500 {
			t.Fatal("Error al calcular el salario, debería ser 1000")
		}
	})
	t.Run("Calcula el salario para categoria B", func(t *testing.T) {
		//Arrange
		var minutos = 60.0
		//Act
		result := calcularSalario(minutos, "B")
		//Assert
		if result != 1500+1500*0.2 {
			t.Fatal("Error al calcular el promedio")
		}
	})
	t.Run("Calcula el salario para categoria C", func(t *testing.T) {
		//Arrange
		var minutos = 60.0
		//Act
		result := calcularSalario(minutos, "C")
		//Assert
		if result != 1000 {
			t.Fatal("Error al calcular el promedio")
		}
	})

}
