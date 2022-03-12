package isValid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	actual := isValid("[]{}()")
	assert.True(t, actual)
}

func TestIsValid2(t *testing.T) {
	actual := isValid("()")
	assert.True(t, actual)
}

func TestIsValid3(t *testing.T) {
	actual := isValid("(]")
	assert.False(t, actual)
}

func TestIsValid4(t *testing.T) {
	actual := isValid("([)]")
	assert.False(t, actual)
}

func TestIsValid5(t *testing.T) {
	actual := isValid("{[]}")
	assert.True(t, actual)
}

func TestIsValid6(t *testing.T) {
	actual := isValid("]")
	assert.False(t, actual)
}
