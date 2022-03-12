package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitIntListNode(nums []int) *ListNode {
	l := &ListNode{}
	for _, num := range nums {
		l.Add(num)
	}

	return l.Next
}

func (head *ListNode) Add(num int) {
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
	}

	tmp.Next = &ListNode{
		Val: num,
	}
}

func (head *ListNode) Equal(other *ListNode) bool {
	for head != nil && other != nil {
		if head.Val != other.Val {
			return false
		}
		head = head.Next
		other = other.Next
	}

	return head == other
}

func (l *ListNode) NextNode() *ListNode {
	if l.Next != nil {
		*l = *l.Next
		return l

	}
	return nil
}
