package searchInsert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	actual := searchInsert([]int{1, 3, 5, 6}, 5)
	expected := 2
	assert.Equal(t, expected, actual)
}

func TestSearchInsert2(t *testing.T) {
	actual := searchInsert([]int{1, 3, 5, 6}, 2)
	expected := 1
	assert.Equal(t, expected, actual)
}

func TestSearchInsert3(t *testing.T) {
	actual := searchInsert([]int{1, 3, 5, 6}, 7)
	expected := 4
	assert.Equal(t, expected, actual)
}

func TestSearchInsert4(t *testing.T) {
	actual := searchInsert([]int{1, 3, 5, 6}, 0)
	expected := 0
	assert.Equal(t, expected, actual)
}

func TestSearchInsert5(t *testing.T) {
	actual := searchInsert([]int{1}, 1)
	expected := 0
	assert.Equal(t, expected, actual)
}
