package mergeKLists

import (
	"testing"

	. "github.com/SnakeHacker/leetcode/algo/common/listnode"
	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	list := []*ListNode{};
	list = append(list, InitIntListNode([]int{1, 4, 5}))
	list = append(list, InitIntListNode([]int{1, 3, 4}))
	list = append(list, InitIntListNode([]int{2,6}))

	actual := mergeKLists(list)
	expected := InitIntListNode([]int{1,1,2,3,4,4,5,6})

	assert.True(t, expected.Equal(actual))
}
