package maxDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := root.Left
	r := root.Right

	if l == nil && r == nil {
		return 1
	}

	ld := maxDepth(l)
	rd := maxDepth(r)
	if ld > rd {
		return 1 + ld
	}

	return 1 + rd
}
