package day_1_calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddition(t *testing.T) {
	assert.Equal(t, addition(1, 2), float64(3))
}

func TestSubtraction(t *testing.T) {
	assert.Equal(t, subtraction(2, 3), float64(-1))
}

func TestMultiplication(t *testing.T) {
	assert.Equal(t, multiplication(2, 3), 6.0)
}

func TestDivision(t *testing.T) {
	assert.Equal(t, division(6, 3), 2.0)
}
