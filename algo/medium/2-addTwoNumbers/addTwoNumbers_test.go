package addTwoNumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := makeListNode([]int{2, 4, 3})
	l2 := makeListNode([]int{5, 6, 4})

	expected := makeListNode([]int{7, 0, 8})
	actual := addTwoNumbers(l1, l2)
	assert.Equal(t, expected, actual)

}

func TestAddTwoNumbers2(t *testing.T) {
	l1 := makeListNode([]int{1, 0, 0, 0, 1})
	l2 := makeListNode([]int{5, 6, 4})

	expected := makeListNode([]int{6, 6, 4, 0, 1})
	actual := addTwoNumbers(l1, l2)

	assert.Equal(t, expected, actual)
}

func TestAddTwoNumbers3(t *testing.T) {
	l1 := &ListNode{Val: 5}
	l2 := &ListNode{Val: 5}

	expected := makeListNode([]int{0, 1})
	actual := addTwoNumbers(l1, l2)

	assert.Equal(t, expected, actual)

}

func TestAddTwoNumbers4(t *testing.T) {
	expected := makeListNode([]int{})
	actual := addTwoNumbers(nil, nil)

	assert.Equal(t, expected, actual)

}

func TestMakeListNode(t *testing.T) {
	l := makeListNode([]int{2, 4, 3})
	assert.Equal(t, 2, l.Val)
	l = l.Next
	assert.Equal(t, 4, l.Val)
	l = l.Next
	assert.Equal(t, 3, l.Val)
	l = l.Next
	assert.Nil(t, l)
}
