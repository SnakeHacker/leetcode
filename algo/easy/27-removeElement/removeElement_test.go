package removeElement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	actual := removeElement([]int{3, 2, 2, 3}, 3)
	expected := 2
	assert.Equal(t, expected, actual)
}

func TestRemoveElement2(t *testing.T) {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	actual := removeElement(arr, 2)
	expected := 5
	assert.Equal(t, expected, actual)
}
