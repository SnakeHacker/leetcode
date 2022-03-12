package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := InitStack()
	s.Push(123)
	s.Push(456)

	actual := s.Pop()
	expected := 456
	assert.Equal(t, expected, actual)

	actual = s.Pop()
	expected = 123
	assert.Equal(t, expected, actual.(int))

	actual = s.Pop()
	assert.Nil(t, actual)

	actual = s.Pop()
	assert.Nil(t, actual)

	s.Push(789)
	actual = s.Pop()
	expected = 789
	assert.Equal(t, expected, actual.(int))

}
