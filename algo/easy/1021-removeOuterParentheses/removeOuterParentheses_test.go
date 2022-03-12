package removeOuterParentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveOuterParentheses(t *testing.T) {
	actual := removeOuterParentheses("(()())(())")
	expected := "()()()"
	assert.Equal(t, expected, actual)
}

func TestRemoveOuterParentheses2(t *testing.T) {
	actual := removeOuterParentheses("(()())(())(()(()))")
	expected := "()()()()(())"
	assert.Equal(t, expected, actual)
}
