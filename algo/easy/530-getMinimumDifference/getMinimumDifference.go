package getMinimumDifference

import (
	"github.com/SnakeHacker/leetcode/algo/common/utils"
)

/*
	给定一个所有节点为非负值的二叉搜索树，求树中任意两节点的差的绝对值的最小值。

示例 :

输入:

   1
    \
     3
    /
   2

输出:
1

解释:
最小绝对差为1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
注意: 树中至少有2个节点。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
@Author: Zhangyang
*/

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMinimumDifference(root *TreeNode) int {
	path := helper(root)
	preValue := path[0]
	result := utils.MAX_INT
	for i := 1; i < len(path); i++ {
		result = min(result, path[i]-preValue)
		preValue = path[i]
	}
	return result
}

func helper(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	return append(append(helper(node.Left), node.Val), helper(node.Right)...)
}
