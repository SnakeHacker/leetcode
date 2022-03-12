package plusOne

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	actual := plusOne([]int{1, 2, 3})
	expected := []int{1, 2, 4}

	assert.Equal(t, actual, expected)
}

func TestPlusOne2(t *testing.T) {
	actual := plusOne([]int{9, 9, 9})
	expected := []int{1, 0, 0, 0}

	assert.Equal(t, actual, expected)
}
