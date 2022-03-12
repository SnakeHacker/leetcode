package arrayPairSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayPairSum(t *testing.T) {
	actual := arrayPairSum([]int{1, 4, 3, 2})
	expected := 4
	assert.Equal(t, expected, actual)
}
