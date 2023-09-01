package ejercicios

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcularAlimento(t *testing.T) {
	t.Run("Calcula el alimento para un perro", func(t *testing.T) {
		//Arrange
		animal := Dog
		cantidad := 10
		//Act
		fx, msg := calcularAlimento(animal)
		result := fx(cantidad)
		//Assert
		assert.Equal(t, result, 100000, "El resultado alimento de perro necesario es de 10000 gr")
		assert.Equal(t, msg, "", "Hay un mensaje de error")
	})
	t.Run("Calcula el alimento para un gato", func(t *testing.T) {
		//Arrange
		animal := Cat
		cantidad := 10
		//Act
		fx, msg := calcularAlimento(animal)
		result := fx(cantidad)
		//Assert
		assert.Equal(t, result, 50000, "El resultado alimento de gato necesario es de 50000 gr")
		assert.Equal(t, msg, "", "Hay un mensaje de error")
	})
	t.Run("Calcula el alimento para un hamster", func(t *testing.T) {
		//Arrange
		animal := Hamster
		cantidad := 10
		//Act
		fx, msg := calcularAlimento(animal)
		result := fx(cantidad)
		//Assert
		assert.Equal(t, result, 2500, "El resultado alimento de hamster necesario es de 2500 gr")
		assert.Equal(t, msg, "", "Hay un mensaje de error")
	})
	t.Run("Calcula el alimento para una tarantula", func(t *testing.T) {
		//Arrange
		animal := Spider
		cantidad := 10
		//Act
		fx, msg := calcularAlimento(animal)
		result := fx(cantidad)
		//Assert
		assert.Equal(t, result, 1500, "El resultado alimento de tarantula necesario es de 2500 gr")
		assert.Equal(t, msg, "", "Hay un mensaje de error")
	})
	t.Run("Error si no existe el animal", func(t *testing.T) {
		//Arrange
		animal := "unicornio"
		//Act
		fx, msg := calcularAlimento(animal)
		//Assert
		assert.Nil(t, fx, "No tiene que existir la funcion", animal)
		assert.Equal(t, msg, "Error", "No hay un mensaje de error")
	})

}
