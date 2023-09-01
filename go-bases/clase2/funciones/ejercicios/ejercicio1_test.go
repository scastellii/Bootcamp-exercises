package ejercicios

import "testing"

func TestCalcularImpuesto(t *testing.T) {
	t.Run("Gana por debajo de 50000", func(t *testing.T) {
		//Arrange
		salario := 49000.0
		//Act
		result := calcularImpuesto(salario)
		//Assert
		if result != 0 {
			t.Fatal("No se le debería cobrar impuesto por cobrar menos de 50000")
		}
	})
	t.Run("Gana por encima de 50000", func(t *testing.T) {
		//Arrange
		salario := 60000.0
		//Act
		result := calcularImpuesto(salario)
		//Assert
		if result != 10200.0 {
			t.Fatal("Se le debería cobrar un 17% de dto")
		}
	})
	t.Run("Gana por encima de 150000", func(t *testing.T) {
		//Arrange
		salario := 160000.0
		//Act
		result := calcularImpuesto(salario)
		//Assert
		if result != salario*0.27 {
			t.Fatal("Se le debería cobrar un 27% de impuesto")
		}
	})
}
