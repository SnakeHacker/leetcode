package merge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	actual := merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	expected := []int{1, 2, 2, 3, 5, 6}
	assert.Equal(t, expected, actual)
}

func TestMerge2(t *testing.T) {
	actual := merge([]int{6, 7, 8, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	expected := []int{2, 5, 6, 6, 7, 8}
	assert.Equal(t, expected, actual)
}

func TestMerge3(t *testing.T) {
	actual := merge([]int{1}, 1, []int{}, 0)
	expected := []int{1}
	assert.Equal(t, expected, actual)
}
