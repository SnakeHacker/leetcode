package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	actual := search(nums, target)
	expected := 4

	assert.Equal(t, expected, actual)
}

func TestSearch2(t *testing.T) {
	nums := []int{-1}
	target := 9
	actual := search(nums, target)
	expected := -1

	assert.Equal(t, expected, actual)
}
