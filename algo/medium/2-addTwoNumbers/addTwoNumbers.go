package addTwoNumbers

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

思路： 一开始想的是先把两个链表分别遍历后相加求和，然后对和制造链表。但是这种方法如果遇到链表长度超过MaxInt就会产生溢出错误。

同时遍历2个链表，每一位相加，设置flag记录进位。

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
@Author: Zhangyang, Mickey
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return helper(l1, l2, 0)
}

func helper(l1 *ListNode, l2 *ListNode, extra int) *ListNode {
	if l1 == nil && l2 == nil && extra == 1 {
		return &ListNode{1, nil}
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	total := l1.Val + l2.Val + extra
	l := &ListNode{total % 10, nil}
	l.Next = helper(l1.Next, l2.Next, total/10)
	return l
}

func makeListNode(nums []int) *ListNode {
	l := &ListNode{}
	for _, num := range nums {
		addListNode(l, num)
	}
	return l.Next
}

func addListNode(head *ListNode, num int) {
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
	}

	tmp.Next = &ListNode{
		Val: num,
	}
}
