package inorderTraversal

import "github.com/SnakeHacker/leetcode/algo/common/tree"

/*
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

/*
@Author: Mickey
*/

func inorderTraversal(root *tree.BinaryTreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	left := inorderTraversal(root.Left)
	for i := range left {
		res = append(res, left[i])
	}

	mid := root.Val
	res = append(res, mid)
	right := inorderTraversal(root.Right)
	for i := range right {
		res = append(res, right[i])
	}
	return res
}
