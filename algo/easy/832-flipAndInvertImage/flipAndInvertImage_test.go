package flipAndInvertImage

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SnakeHacker/leetcode/algo/common/array"
)

func TestFlipAndInvertImage(t *testing.T) {
	nums := []int{0, 0, 0, 0, 1, 1, 1, 1}
	row := 2
	column := 4
	matrix := array.Init2DIntArray(nums, row, column)
	acutal := flipAndInvertImage(matrix)
	expected := [][]int{[]int{1, 1, 1, 1}, []int{0, 0, 0, 0}}
	assert.Equal(t, expected, acutal)
}

func TestFlipAndInvertImage2(t *testing.T) {
	nums := []int{1, 1, 0, 1, 0, 1, 0, 0, 0}
	row := 3
	column := 3
	matrix := array.Init2DIntArray(nums, row, column)
	acutal := flipAndInvertImage(matrix)
	expected := [][]int{[]int{1, 0, 0}, []int{0, 1, 0}, []int{1, 1, 1}}
	assert.Equal(t, expected, acutal)
}
