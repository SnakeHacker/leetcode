package generate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	actual := generate(5)
	expected := [][]int{[]int{1}, []int{1, 1}, []int{1, 2, 1}, []int{1, 3, 3, 1}, []int{1, 4, 6, 4, 1}}
	assert.Equal(t, actual, expected)
}

func TestGenerate2(t *testing.T) {
	actual := generate(0)
	expected := [][]int{}
	assert.Equal(t, actual, expected)
}

func TestGenerate3(t *testing.T) {
	actual := generate(3)
	expected := [][]int{[]int{1}, []int{1, 1}, []int{1, 2, 1}}
	assert.Equal(t, actual, expected)
}
