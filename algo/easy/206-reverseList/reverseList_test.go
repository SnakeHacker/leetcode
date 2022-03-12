package reverseList

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/SnakeHacker/leetcode/algo/common/listnode"
)

func TestReverseList(t *testing.T) {
	actual := reverseList(InitIntListNode([]int{1, 2, 3, 4}))
	expected := InitIntListNode([]int{4, 3, 2, 1})
	assert.Equal(t, expected, actual)
}
