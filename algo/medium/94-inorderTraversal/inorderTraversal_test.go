package inorderTraversal

import (
	"testing"

	"github.com/SnakeHacker/leetcode/algo/common/tree"
	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal(t *testing.T) {
	root := &tree.BinaryTreeNode{
		Val: 1,
		Right: &tree.BinaryTreeNode{
			Val: 2,
			Left: &tree.BinaryTreeNode{
				Val: 3,
			},
		},
	}

	actual := inorderTraversal(root)
	expected := []int{1, 3, 2}
	assert.Equal(t, expected, actual)
}
