package lengthOfLongestSubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	actual := lengthOfLongestSubstring("abcabcbb")
	expected := 3
	assert.Equal(t, expected, actual)
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	actual := lengthOfLongestSubstring("bbbbb")
	expected := 1
	assert.Equal(t, expected, actual)
}

func TestLengthOfLongestSubstring3(t *testing.T) {
	actual := lengthOfLongestSubstring("pwwkew")
	expected := 3
	assert.Equal(t, expected, actual)
}

func TestLengthOfLongestSubstring4(t *testing.T) {
	actual := lengthOfLongestSubstring("")
	expected := 0
	assert.Equal(t, expected, actual)
}

func TestLengthOfLongestSubstring5(t *testing.T) {
	actual := lengthOfLongestSubstring(" ")
	expected := 1
	assert.Equal(t, expected, actual)
}

func TestLengthOfLongestSubstring6(t *testing.T) {
	actual := lengthOfLongestSubstring("au")
	expected := 2
	assert.Equal(t, expected, actual)
}

func TestLengthOfLongestSubstring7(t *testing.T) {
	actual := lengthOfLongestSubstring("abba")
	expected := 2
	assert.Equal(t, expected, actual)
}
