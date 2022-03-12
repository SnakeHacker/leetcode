package listnode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitIntListNode(t *testing.T) {
	root := InitIntListNode([]int{2, 3, 4, 5})
	expected := 2
	assert.Equal(t, expected, root.Val)

	next := root.NextNode()
	actual := next.Val
	expected = 3
	assert.Equal(t, expected, actual)

	next = next.NextNode()
	actual = next.Val
	expected = 4
	assert.Equal(t, expected, actual)

	next = next.NextNode()
	actual = next.Val
	expected = 5
	assert.Equal(t, expected, actual)

	next = next.NextNode()
	assert.Nil(t, next)
}

func TestEqualIntListNode(t *testing.T) {
	root1 := InitIntListNode([]int{2, 3, 4, 5})
	root2 := InitIntListNode([]int{2, 3, 4, 5})
	root3 := InitIntListNode([]int{2, 3, 4})

	assert.True(t, root1.Equal(root2))
	assert.False(t, root1.Equal(root3))
}
