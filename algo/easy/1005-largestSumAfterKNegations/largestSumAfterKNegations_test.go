package largestSumAfterKNegations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestSumAfterKNegations(t *testing.T) {
	actual := largestSumAfterKNegations([]int{4, 2, 3}, 1)
	expected := 5
	assert.Equal(t, expected, actual)
}

func TestLargestSumAfterKNegations2(t *testing.T) {
	actual := largestSumAfterKNegations([]int{3, -1, 0, 2}, 3)
	expected := 6
	assert.Equal(t, expected, actual)
}

func TestLargestSumAfterKNegations3(t *testing.T) {
	actual := largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2)
	expected := 13
	assert.Equal(t, expected, actual)
}
