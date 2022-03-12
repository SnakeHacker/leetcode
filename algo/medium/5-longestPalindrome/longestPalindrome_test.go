package longestPalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	actual := longestPalindrome("aaa")
	expected := "aaa"

	assert.Equal(t, expected, actual)
}

func TestLongestPalindrome2(t *testing.T) {
	actual := longestPalindrome("babad")
	expected := "bab"

	assert.Equal(t, expected, actual)
}

func TestLongestPalindrome3(t *testing.T) {
	actual := longestPalindrome("a")
	expected := "a"

	assert.Equal(t, expected, actual)
}
