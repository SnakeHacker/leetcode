package quickSort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	arr := []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9}
	expected := []int{-11, 1, 2, 4, 6, 8, 9, 22, 33, 34, 44, 54, 55}
	assert.Equal(t, expected, quickSort(&arr, 0, len(arr)-1))
}
