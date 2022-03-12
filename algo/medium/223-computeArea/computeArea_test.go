package computeArea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeArea(t *testing.T) {
	actual := computeArea(-3, 0, 3, 4, 0, -1, 9, 2)
	expected := 45
	assert.Equal(t, expected, actual)
}

func TestComputeArea2(t *testing.T) {
	actual := computeArea(-2, -2, 2, 2, 3, 3, 4, 4)
	expected := 17
	assert.Equal(t, expected, actual)
}
