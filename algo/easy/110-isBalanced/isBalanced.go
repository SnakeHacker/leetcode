package isBalanced

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helper(root) != -1
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := helper(root.Left)
	right := helper(root.Right)
	maxDepth := int(math.Abs(float64(left - right)))
	if left == -1 || right == -1 || maxDepth > 1 {
		return -1
	}

	return int(math.Max(float64(left), float64(right))) + 1
}
