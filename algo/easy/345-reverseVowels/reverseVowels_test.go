package reverseVowels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseVowels(t *testing.T) {
	actual := reverseVowels("hello")
	expected := "holle"
	assert.Equal(t, expected, actual)
}

func TestReverseVowels2(t *testing.T) {
	actual := reverseVowels("leetcode")
	expected := "leotcede"
	assert.Equal(t, expected, actual)
}
