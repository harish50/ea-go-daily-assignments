package day1

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

func TestSin(t *testing.T) {
	assert.Equal(t, sin(0), 0.0)
}

func TestCos(t *testing.T) {
	assert.Equal(t, cos(0), 1.0)
}

func TestTan(t *testing.T) {
	assert.Equal(t, tan(0), 0.0)
}

func TestSqrt(t *testing.T) {
	assert.Equal(t, sqrt(4), 2.0)
}
