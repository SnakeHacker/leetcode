package rob

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob(t *testing.T) {
	actual := rob([]int{1, 2, 3, 1})
	expected := 4
	assert.Equal(t, expected, actual)
}

func TestRob2(t *testing.T) {
	actual := rob([]int{2, 7, 9, 3, 1})
	expected := 12
	assert.Equal(t, expected, actual)
}
