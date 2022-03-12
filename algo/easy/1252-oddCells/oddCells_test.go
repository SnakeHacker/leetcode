package oddCells

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddCells(t *testing.T) {
	arr := [][]int{[]int{0, 1}, []int{1, 1}}
	actual := oddCells(2, 3, arr)
	assert.Equal(t, 6, actual)

	arr = [][]int{[]int{1, 1}, []int{1, 1}}
	actual = oddCells(2, 2, arr)
	assert.Equal(t, 0, actual)

	arr = [][]int{[]int{1, 1}, []int{0, 0}}
	actual = oddCells(2, 2, arr)
	assert.Equal(t, 0, actual)

	arr = [][]int{[]int{40, 5}}
	actual = oddCells(48, 37, arr)
	assert.Equal(t, 83, actual)
}
