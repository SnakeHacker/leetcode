package isPalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome0(t *testing.T) {
	expect := true
	assert.Equal(t, expect, isPalindrome(121))
	assert.Equal(t, expect, isPalindromeV1(121))
}

func TestIsPalindrome1(t *testing.T) {
	expect := false
	assert.Equal(t, expect, isPalindrome(-121))
	assert.Equal(t, expect, isPalindrome(-121))
}

func TestIsPalindrome2(t *testing.T) {
	expect := false
	assert.Equal(t, expect, isPalindrome(10))
	assert.Equal(t, expect, isPalindrome(10))
}

func TestIsPalindrome3(t *testing.T) {
	expect := false
	assert.Equal(t, expect, isPalindrome(1000021))
	assert.Equal(t, expect, isPalindrome(1000021))
}
