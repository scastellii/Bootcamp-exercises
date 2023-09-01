package ejercicios

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRealizarOpreacion(t *testing.T) {
	t.Run("Calcula el Maximo de notas de un alumno", func(t *testing.T) {
		//Arrange
		operacion := Maximum
		notas := []int{5, 7, 8, 9, 1}
		//Act
		fx, msg := realizarOperacion(operacion)
		result := fx(notas...)
		//Assert
		assert.Equal(t, result, 9.0, "El resultado maximo es 9")
		assert.Equal(t, msg, "", "Hay un mensaje de error")
	})
	t.Run("Calcula el minimo de notas de un alumno", func(t *testing.T) {
		//Arrange
		operacion := Minimum
		notas := []int{5, 7, 8, 9, 1}
		//Act
		fx, msg := realizarOperacion(operacion)
		result := fx(notas...)
		//Assert
		assert.Equal(t, result, 1.0, "El resultado minimo es 1")
		assert.Equal(t, msg, "", "Hay un mensaje de error")
	})
	t.Run("Calcula el promedio de notas de un alumno", func(t *testing.T) {
		//Arrange
		operacion := Average
		notas := []int{5, 7, 8, 9, 1}
		//Act
		fx, msg := realizarOperacion(operacion)
		result := fx(notas...)
		//Assert
		assert.Equal(t, result, 6.0, "El promedio es 9")
		assert.Equal(t, msg, "", "Hay un mensaje de error")
	})
	t.Run("Falla si se envía una operación invalida", func(t *testing.T) {
		//Arrange
		operacion := "HOLA"
		//Act
		fx, msg := realizarOperacion(operacion)
		//Assert
		assert.Nil(t, fx, "No debe existir una función HOLA")
		assert.Equal(t, msg, "Error", "No hay un mensaje de error")
	})

}
