package isSubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	actual := isSubsequence("abc", "ahbgdc")
	expected := true
	assert.Equal(t, actual, expected)
}

func TestIsSubsequence2(t *testing.T) {
	actual := isSubsequence("axc", "ahbgdc")
	expected := false
	assert.Equal(t, actual, expected)
}

func TestIsSubsequence3(t *testing.T) {
	actual := isSubsequence("", "ahbgdc")
	expected := true
	assert.Equal(t, actual, expected)
}
