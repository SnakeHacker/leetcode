package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	actual := game([]int{1, 2, 3}, []int{1, 2, 3})
	assert.Equal(t, 3, actual)

	actual = game([]int{2, 2, 3}, []int{3, 2, 1})
	assert.Equal(t, 1, actual)
}
