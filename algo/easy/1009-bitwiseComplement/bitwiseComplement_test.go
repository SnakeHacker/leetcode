package bitwiseComplement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitwiseComplement(t *testing.T) {
	actual := bitwiseComplement(5)
	expected := 2
	assert.Equal(t, expected, actual)

	actual = bitwiseComplementV1(5)
	assert.Equal(t, expected, actual)

}

func TestBitwiseComplement2(t *testing.T) {
	actual := bitwiseComplement(10)
	expected := 5
	assert.Equal(t, expected, actual)

	actual = bitwiseComplementV1(10)
	assert.Equal(t, expected, actual)
}

func TestBitwiseComplement3(t *testing.T) {
	actual := bitwiseComplement(0)
	expected := 1
	assert.Equal(t, expected, actual)

	actual = bitwiseComplementV1(0)
	assert.Equal(t, expected, actual)
}

func TestBitwiseComplement4(t *testing.T) {
	actual := bitwiseComplement(1)
	expected := 0
	assert.Equal(t, expected, actual)

	actual = bitwiseComplementV1(1)
	assert.Equal(t, expected, actual)
}
