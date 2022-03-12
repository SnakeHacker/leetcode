package intersection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	actual := intersection([]int{1, 2, 2, 1}, []int{2, 2})
	expected := []int{2}
	assert.Equal(t, expected, actual)

	actual = intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	expected = []int{9, 4}
	assert.Equal(t, expected, actual)
}
