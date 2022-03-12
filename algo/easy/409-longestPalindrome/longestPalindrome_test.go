package longestPalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	actual := longestPalindrome("abccccdd")
	expected := 7

	assert.Equal(t, expected, actual)
}
