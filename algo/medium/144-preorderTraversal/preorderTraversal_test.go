package preorderTraversal

import (
	"testing"

	"github.com/SnakeHacker/leetcode/algo/common/tree"
	"github.com/stretchr/testify/assert"
)

func TestPreorderTraversal(t *testing.T) {
	root := &tree.BinaryTreeNode{
		Val: 1,
		Right: &tree.BinaryTreeNode{
			Val: 2,
			Left: &tree.BinaryTreeNode{
				Val: 3,
			},
		},
	}

	actual := preorderTraversal(root)
	expected := []int{1, 2, 3}
	assert.Equal(t, expected, actual)
}

func TestPreorderTraversal2(t *testing.T) {
	root := &tree.BinaryTreeNode{
		Val: 1,
		Right: &tree.BinaryTreeNode{
			Val: 2,
			Left: &tree.BinaryTreeNode{
				Val: 3,
			},
		},
	}

	actual := preorderTraversal2(root)
	expected := []int{1, 2, 3}
	assert.Equal(t, expected, actual)
}

func TestPreorderTraversal3(t *testing.T) {
	root := &tree.BinaryTreeNode{
		Val: 3,
		Left: &tree.BinaryTreeNode{
			Val: 1,
		},
		Right: &tree.BinaryTreeNode{
			Val: 2,
		},
	}

	actual := preorderTraversal2(root)
	expected := []int{3, 1, 2}
	assert.Equal(t, expected, actual)
}
