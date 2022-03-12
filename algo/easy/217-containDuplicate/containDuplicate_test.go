package containDuplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	actual := containsDuplicate([]int{1, 2, 3, 4})
	expected := false
	assert.Equal(t, expected, actual)
}

func TestContainsDuplicate2(t *testing.T) {
	actual := containsDuplicate([]int{1, 2, 3, 1})
	expected := true
	assert.Equal(t, expected, actual)
}
