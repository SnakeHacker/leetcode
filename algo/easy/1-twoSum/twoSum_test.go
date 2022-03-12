package twoSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	expected := []int{0, 1}
	actual := twoSum(nums, target)

	assert.Equal(t, expected, actual)
}

func TestTwoSum2(t *testing.T) {
	nums := []int{2, -1}
	target := 9

	expected := []int{}
	actual := twoSum(nums, target)

	assert.Equal(t, expected, actual)
}
