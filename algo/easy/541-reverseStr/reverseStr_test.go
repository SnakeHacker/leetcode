package reverseStr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseStr(t *testing.T) {
	actual := reverseStr("abcdefg", 2)
	expected := "bacdfeg"
	assert.Equal(t, expected, actual)
}

func TestReverseStr2(t *testing.T) {
	actual := reverseStr("abcdefg", 3)
	expected := "cbadefg"
	assert.Equal(t, expected, actual)
}

func TestReverseStr3(t *testing.T) {
	actual := reverseStr("abcdefg", 8)
	expected := "gfedcba"
	assert.Equal(t, expected, actual)
}

func TestReverseStr4(t *testing.T) {
	actual := reverseStr("a", 2)
	expected := "a"
	assert.Equal(t, expected, actual)
}
