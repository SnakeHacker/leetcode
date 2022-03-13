package moveZeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	actual := moveZeroes([]int{0, 1, 0, 3, 12})
	expected := []int{1, 3, 12, 0, 0}
	assert.Equal(t, expected, actual)
}

func TestMoveZeroes2(t *testing.T) {
	actual := moveZeroes([]int{1, 12, 0, 3, 12, 0})
	expected := []int{1, 12, 3, 12, 0, 0}
	assert.Equal(t, expected, actual)
}
