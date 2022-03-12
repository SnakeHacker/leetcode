package mergeKLists

import . "github.com/SnakeHacker/leetcode/algo/common/listnode"

/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/

/*
@Author: Yang
*/

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)

	if length == 0 {
		return nil
	}

	if length == 1 {
		return lists[0]
	}

	return mergeTwoLists(mergeKLists(lists[:length/2]), mergeKLists(lists[length/2:]))
}

// 21-mergeTwoLists
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	p := dummyNode

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			p = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			p = l2
			l2 = l2.Next
		}
	}

	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}

	return dummyNode.Next
}
