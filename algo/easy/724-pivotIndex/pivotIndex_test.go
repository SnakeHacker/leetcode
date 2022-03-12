package pivotIndex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotIndex(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	actual := pivotIndex(nums)
	expected := 3
	assert.Equal(t, expected, actual)
}

func TestPivotIndex2(t *testing.T) {
	nums := []int{1, 2, 3}
	actual := pivotIndex(nums)
	expected := -1
	assert.Equal(t, expected, actual)
}
