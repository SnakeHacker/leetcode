package removeDuplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	actual := removeDuplicates([]int{1, 1, 2})
	expected := 2
	assert.Equal(t, expected, actual)
}

func TestRemoveDuplicates2(t *testing.T) {
	actual := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	expected := 5
	assert.Equal(t, expected, actual)
}
