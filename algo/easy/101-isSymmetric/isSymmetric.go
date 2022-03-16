package isSymmetric

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := root.Left
	right := root.Right
	return equal(left, right)
}

func equal(left, right *TreeNode) bool {
	if left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		return equal(left.Left, right.Right) && equal(left.Right, right.Left)
	}

	if left == nil && right == nil {
		return true
	}

	return false
}
