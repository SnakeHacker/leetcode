package mergeTwoLists

import (
	"testing"

	. "github.com/SnakeHacker/leetcode/algo/common/listnode"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	l1 := InitIntListNode([]int{1, 2, 4})
	l2 := InitIntListNode([]int{1, 3, 4})

	actual := mergeTwoLists(l1, l2)
	expected := InitIntListNode([]int{1, 1, 2, 3, 4, 4})

	assert.True(t, expected.Equal(actual))
}
