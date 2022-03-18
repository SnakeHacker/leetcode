package popSort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopSort(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5}, popSort([]int{2, 3, 4, 5, 1}))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, popSort([]int{2, 4, 1, 5, 3}))

}
