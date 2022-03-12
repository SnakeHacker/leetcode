package reverseList

import (
	. "github.com/SnakeHacker/leetcode/algo/common/listnode"
)

/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
*/

/*
@Author: Mickey
*/

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := *head.Next
	head.Next = nil
	return helper(&next, head)
}

func helper(node *ListNode, pre *ListNode) *ListNode {
	tmp := *node
	if tmp.Next == nil {
		tmp.Next = pre
		return &tmp
	}
	tmp.Next = pre
	return helper(node.Next, &tmp)
}
