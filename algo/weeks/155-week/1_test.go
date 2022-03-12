package week

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumPrimeArrangements(t *testing.T) {
	actual := numPrimeArrangements(5)
	expected := 12
	assert.Equal(t, expected, actual)
}

func TestNumPrimeArrangements2(t *testing.T) {
	actual := numPrimeArrangements(100)
	expected := 682289015
	assert.Equal(t, expected, actual)
}

func TestNumPrimeArrangements3(t *testing.T) {
	actual := numPrimeArrangements(7)
	expected := int(factorial(3) * factorial(4))
	assert.Equal(t, expected, actual)
}
