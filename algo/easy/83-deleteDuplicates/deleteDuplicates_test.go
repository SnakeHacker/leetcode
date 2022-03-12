package deleteDuplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/SnakeHacker/leetcode/algo/common/listnode"
)

func TestDeleteDuplicates(t *testing.T) {
	l := InitIntListNode([]int{1, 1, 2})
	actual := deleteDuplicates(l)
	expected := InitIntListNode([]int{1, 2})
	assert.Equal(t, expected, actual)
}

func TestDeleteDuplicates2(t *testing.T) {
	l := InitIntListNode([]int{1, 1, 2, 3, 3})
	actual := deleteDuplicates(l)
	expected := InitIntListNode([]int{1, 2, 3})
	assert.Equal(t, expected, actual)
}

func TestDeleteDuplicates3(t *testing.T) {
	l := InitIntListNode([]int{1, 1, 1})
	actual := deleteDuplicates(l)
	expected := InitIntListNode([]int{1})
	assert.Equal(t, expected, actual)
}
