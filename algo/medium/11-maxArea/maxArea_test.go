package maxArea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	actual := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	expected := 49
	assert.Equal(t, expected, actual)
}
