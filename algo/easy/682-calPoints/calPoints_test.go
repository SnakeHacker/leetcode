package calPoints

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalPoints(t *testing.T) {
	actual := calPoints([]string{"5", "2", "C", "D", "+"})
	expected := 30
	assert.Equal(t, expected, actual)
}

func TestCalPoints2(t *testing.T) {
	actual := calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"})
	expected := 27
	assert.Equal(t, expected, actual)
}
