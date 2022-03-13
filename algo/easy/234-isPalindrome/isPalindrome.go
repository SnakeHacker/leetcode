package isPalindrome

import (
	. "github.com/SnakeHacker/leetcode/algo/common/listnode"
)

/**
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。



示例 1：


输入：head = [1,2,2,1]
输出：true
示例 2：


输入：head = [1,2]
输出：false


提示：

链表中节点数目在范围[1, 105] 内
0 <= Node.val <= 9


进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

*/
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			break
		} else {
			fast = fast.Next
		}
	}

	if fast != nil {
		slow = slow.Next
	}

	l1 := head
	l2 := reverseList(slow)

	for l2 != nil {
		if l1.Val == l2.Val {
			l1 = l1.Next
			l2 = l2.Next
		} else {
			return false
		}
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var preNode *ListNode

	currNode := head

	for currNode != nil {
		next := currNode.Next
		currNode.Next = preNode
		preNode = currNode
		currNode = next
	}

	return preNode
}
