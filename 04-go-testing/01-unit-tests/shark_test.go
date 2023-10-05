package hunt

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSharkHuntsSuccessfully(t *testing.T) {
	// Arrange
	shark := Shark{hungry: true, tired: false, speed: 12}
	prey := Prey{name: "fish", speed: 11}
	/*
		sharkFull := Shark{hungry: false, tired: true, speed: 12}*/
	// Act
	_ = shark.Hunt(&prey)
	// Assert
	assert.True(t, shark.tired)
	assert.False(t, shark.hungry)
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	// Arrange
	shark := Shark{hungry: true, tired: true, speed: 10}
	prey := Prey{name: "fish", speed: 5}
	var err error
	msg := "cannot hunt, i am really tired"
	err = errors.New(msg)
	// Act
	result := shark.Hunt(&prey)
	// Assert
	assert.Equal(t, err, result)
}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
	// Arrange
	shark := Shark{hungry: false, tired: false, speed: 10}
	prey := Prey{name: "fish", speed: 5}
	var err error
	msg := "cannot hunt, i am not hungry"
	err = errors.New(msg)
	// Act
	result := shark.Hunt(&prey)
	// Assert
	assert.Equal(t, err, result)
}

func TestSharkCannotReachThePrey(t *testing.T) {
	// Arrange
	shark := Shark{hungry: true, tired: false, speed: 10}
	prey := Prey{name: "fish", speed: 11}
	msg := "could not catch it"
	// Act
	result := shark.Hunt(&prey)
	// Assert
	assert.EqualError(t, result, msg)
}

func TestSharkHuntNilPrey(t *testing.T) {
	// Arrange
	shark := Shark{hungry: true, tired: false, speed: 10}
	var prey *Prey
	msg := "Prey is null!"
	// Act
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Prey is null!")
		}
	}()
	result := shark.Hunt(prey)

	// Assert
	defer assert.Error(t, result, msg)
}
