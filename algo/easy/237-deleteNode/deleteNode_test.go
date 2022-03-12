package deleteNode

import (
	"testing"

	. "github.com/SnakeHacker/leetcode/algo/common/listnode"
	"github.com/stretchr/testify/assert"
)

func TestDeleteNode(t *testing.T) {
	actual := InitIntListNode([]int{4, 5, 1, 9})
	next := actual.Next
	deleteNode(next)
	expected := InitIntListNode([]int{4, 1, 9})
	assert.Equal(t, expected, actual)
}
