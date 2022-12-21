package day_1_calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddition(t *testing.T) {
	assert.Equal(t, addition(1, 2), float64(3))
}
