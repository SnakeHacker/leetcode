package maxSubArray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	actual := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	expected := 6
	assert.Equal(t, expected, actual)
}
