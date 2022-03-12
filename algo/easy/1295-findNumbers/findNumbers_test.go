package findNumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindNumbers(t *testing.T) {
	actual := findNumbers([]int{12, 345, 2, 6, 7896})
	expected := 2
	assert.Equal(t, expected, actual)
}
