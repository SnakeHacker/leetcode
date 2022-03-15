package postorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return []int{}
	}

	left := postorderTraversal(root.Left)
	right := postorderTraversal(root.Right)

	res = append(res, left...)
	res = append(res, right...)
	res = append(res, root.Val)

	return res
}
