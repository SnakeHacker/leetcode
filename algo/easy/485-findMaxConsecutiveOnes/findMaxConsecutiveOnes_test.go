package findMaxConsecutiveOnes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	expected := 3
	actual := findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})

	assert.Equal(t, expected, actual)
}
