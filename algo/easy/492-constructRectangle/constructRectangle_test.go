package constructRectangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructRectangle(t *testing.T) {
	actual := constructRectangle(4)
	expected := []int{2, 2}
	assert.Equal(t, expected, actual)
}

func TestConstructRectangle2(t *testing.T) {
	actual := constructRectangle(1)
	expected := []int{1, 1}
	assert.Equal(t, expected, actual)
}

func TestConstructRectangle3(t *testing.T) {
	actual := constructRectangle(2)
	expected := []int{2, 1}
	assert.Equal(t, expected, actual)
}
