package diStringMatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiStringMatch(t *testing.T) {
	actual := diStringMatch("IDID")
	expected := []int{0, 4, 1, 3, 2}
	assert.Equal(t, expected, actual)
}

func TestDiStringMatch2(t *testing.T) {
	actual := diStringMatch("III")
	expected := []int{0, 1, 2, 3}
	assert.Equal(t, expected, actual)
}

func TestDiStringMatch3(t *testing.T) {
	actual := diStringMatch("DDI")
	expected := []int{3, 2, 0, 1}
	assert.Equal(t, expected, actual)
}
