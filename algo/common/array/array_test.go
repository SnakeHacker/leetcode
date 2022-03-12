package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit2DIntArray(t *testing.T) {
	actual := Init2DIntArray([]int{2, 3, 4, 5, 6}, 2, 3)
	expected := [][]int{[]int{2, 3, 4}, []int{5, 6}}
	assert.Equal(t, expected, actual)
}

func TestInit2DIntArray2(t *testing.T) {
	actual := Init2DIntArray([]int{1, 3, 5, 7, 9, 11}, 2, 3)
	expected := [][]int{[]int{1, 3, 5}, []int{7, 9, 11}}
	assert.Equal(t, expected, actual)
}
