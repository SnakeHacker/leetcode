package isRectangleOverlap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRectangleOverlap(t *testing.T) {
	actual := isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3})
	assert.True(t, actual)
}

func TestIsRectangleOverlap2(t *testing.T) {
	actual := isRectangleOverlap([]int{0, 0, 1, 1}, []int{1, 0, 2, 1})
	assert.False(t, actual)
}
