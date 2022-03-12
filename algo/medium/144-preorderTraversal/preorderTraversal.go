package preorderTraversal

import (
	"github.com/SnakeHacker/leetcode/algo/common/stack"
	"github.com/SnakeHacker/leetcode/algo/common/tree"
)

/*
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

/*
@Author: Mickey
*/

func preorderTraversal(root *tree.BinaryTreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	mid := root.Val
	res = append(res, mid)

	left := preorderTraversal(root.Left)
	for i := range left {
		res = append(res, left[i])
	}

	right := preorderTraversal(root.Right)
	for i := range right {
		res = append(res, right[i])
	}
	return res
}

func preorderTraversal2(root *tree.BinaryTreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	stack := stack.InitStack()
	stack.Push(root)

	for {
		node := stack.Pop()
		if node == nil {
			break
		}
		res = append(res, node.(*tree.BinaryTreeNode).Val)

		right := node.(*tree.BinaryTreeNode).Right
		if right != nil {
			stack.Push(right)
		}

		left := node.(*tree.BinaryTreeNode).Left
		if left != nil {
			stack.Push(left)
		}
	}

	return res
}
