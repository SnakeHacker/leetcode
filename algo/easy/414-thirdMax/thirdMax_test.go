package thirdMax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThirdMax(t *testing.T) {
	actual := thirdMax([]int{3, 2, 1})
	expected := 1
	assert.Equal(t, expected, actual)
}

func TestThirdMax2(t *testing.T) {
	actual := thirdMax([]int{1, 2})
	expected := 2
	assert.Equal(t, expected, actual)
}

func TestThirdMax3(t *testing.T) {
	actual := thirdMax([]int{2, 2, 3, 1})
	expected := 1
	assert.Equal(t, expected, actual)
}

func TestThirdMax4(t *testing.T) {
	actual := thirdMax([]int{5, 2, 2})
	expected := 5
	assert.Equal(t, expected, actual)
}

func TestThirdMax5(t *testing.T) {
	actual := thirdMax([]int{1, 2, -2147483648})
	expected := -2147483648
	assert.Equal(t, expected, actual)
}

func TestThirdMax6(t *testing.T) {
	actual := thirdMax([]int{1, 2, 2, 5, 3, 5})
	expected := 2
	assert.Equal(t, expected, actual)
}
