package findDisappearedNumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDisappearedNumbers(t *testing.T) {
	actual := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	expected := []int{5, 6}
	assert.Equal(t, expected, actual)
}
