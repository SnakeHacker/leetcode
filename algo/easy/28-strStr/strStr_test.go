package strStr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	actual := strStr("hello", "ll")
	expected := 2
	assert.Equal(t, expected, actual)
}

func TestStrStr2(t *testing.T) {
	actual := strStr("aaaaa", "bba")
	expected := -1
	assert.Equal(t, expected, actual)
}

func TestStrStr3(t *testing.T) {
	actual := strStr("aaaaa", "aaab")
	expected := -1
	assert.Equal(t, expected, actual)
}

func TestStrStr4(t *testing.T) {
	actual := strStr("", "aaab")
	expected := -1
	assert.Equal(t, expected, actual)
}

func TestStrStr5(t *testing.T) {
	actual := strStr("", "")
	expected := 0
	assert.Equal(t, expected, actual)
}

func TestStrStr6(t *testing.T) {
	actual := strStr("", "sd")
	expected := -1
	assert.Equal(t, expected, actual)
}

func TestStrStr7(t *testing.T) {
	actual := strStr("aaa", "aaa")
	expected := 0
	assert.Equal(t, expected, actual)
}

func TestStrStr8(t *testing.T) {
	actual := strStr("mississippi", "issip")
	expected := 4
	assert.Equal(t, expected, actual)
}
