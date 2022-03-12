package getMinimumDifference

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func parseTree(originalS string) *TreeNode {
	s := originalS[1 : len(originalS)-1]
	var r []*TreeNode
	for _, numS := range strings.Split(s, ",") {
		if numS == "null" {
			r = append(r, nil)
		} else {
			num, _ := strconv.Atoi(numS)
			r = append(r, &TreeNode{Val: num})
		}
	}
	return initBinaryTree(r)
}

// https://support.leetcode.com/hc/en-us/articles/360011883654-What-does-1-null-2-3-mean-in-binary-tree-representation-
func initBinaryTree(treeArray []*TreeNode) *TreeNode {
	length := len(treeArray)
	if length == 0 {
		return nil
	}

	var preTrees []*TreeNode
	root := treeArray[0]
	preTrees = []*TreeNode{root}

	for depth := 1; ; depth++ {
		var tPreTrees []*TreeNode
		t := 1 << uint(depth)
		if t > length {
			break
		}

		curLevel := treeArray[t-1 : min(t+1+t>>1, length)]

		for i, tree := range preTrees {
			li, ri := 2*i, 2*i+1
			if tree != nil && li < len(curLevel) {
				tree.Left = curLevel[li]
				tPreTrees = append(tPreTrees, tree.Left)

			}
			if tree != nil && ri < len(curLevel) {
				tree.Right = curLevel[ri]
				tPreTrees = append(tPreTrees, tree.Right)
			}
			firstNotNillIndex := len(tPreTrees)

			for i, tree := range tPreTrees {
				if tree != nil {
					firstNotNillIndex = i
					break
				}
			}
			preTrees = tPreTrees[firstNotNillIndex:]
		}
	}

	return root
}

func TestDiff1(t *testing.T) {
	expected := 1
	actual := getMinimumDifference(parseTree("[1,null,3,2]"))

	assert.Equal(t, expected, actual)
}

func TestDiff2(t *testing.T) {
	expected := 9
	actual := getMinimumDifference(parseTree("[236,104,701,null,227,null,911]"))

	assert.Equal(t, expected, actual)
}

func TestDiff3(t *testing.T) {
	expected := 1
	actual := getMinimumDifference(parseTree("[3,0,5,null,null,4,8]"))

	assert.Equal(t, expected, actual)
}

func TestDiff4(t *testing.T) {
	expected := 1
	actual := getMinimumDifference(parseTree("[3,1,6,0,null,4,8]"))

	assert.Equal(t, expected, actual)
}
