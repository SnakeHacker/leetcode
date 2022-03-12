package sortedSquares

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	actual := sortedSquares([]int{-4, -1, 0, 3, 10})
	expected := []int{0, 1, 9, 16, 100}
	assert.Equal(t, expected, actual)
}

func TestSortedSquares2(t *testing.T) {
	actual := sortedSquares2([]int{-4, -1, 0, 3, 10})
	expected := []int{0, 1, 9, 16, 100}
	assert.Equal(t, expected, actual)
}
