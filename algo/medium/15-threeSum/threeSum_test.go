package threeSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	actual := threeSum([]int{-1, 0, 1, 2, -1, -4})
	expected := [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}
	assert.Equal(t, expected, actual)
}

func TestThreeSum2(t *testing.T) {
	actual := threeSum([]int{-1, 0, 0, 0, 1, -4})
	expected := [][]int{[]int{-1, 0, 1}, []int{0, 0, 0}}
	assert.Equal(t, expected, actual)
}
