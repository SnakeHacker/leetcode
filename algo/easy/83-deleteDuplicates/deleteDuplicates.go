package deleteDuplicates

import (
	. "github.com/SnakeHacker/leetcode/algo/common/listnode"
)

/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
*/

/*
@Author: Mickey
*/

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	tmp := head
	for tmp != nil {
		if tmp.Next == nil {
			break
		}
		for tmp.Val == tmp.Next.Val {
			*tmp = *tmp.Next
			if tmp.Next == nil {
				break
			}
		}
		tmp = tmp.Next
	}

	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	tmp := head
	for tmp != nil {
		if tmp.Next == nil {
			break
		}
		if tmp.Val == tmp.Next.Val {
			*tmp = *tmp.Next
			if tmp.Next == nil {
				break
			}
			continue
		}
		tmp = tmp.Next
	}

	return head
}
