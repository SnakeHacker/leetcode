package threeSumClosest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSumClosest(t *testing.T) {
	actual := threeSumClosest([]int{-1, 2, 1, -4}, 1)
	expected := 2
	assert.Equal(t, expected, actual)
}
