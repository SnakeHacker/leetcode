package backspaceCompare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackspaceCompare(t *testing.T) {
	assert.True(t, backspaceCompare("ab#c", "ad#c"))
	assert.True(t, backspaceCompare("ab##", "c#d#"))
	assert.True(t, backspaceCompare("a##c", "#a#c"))
	assert.False(t, backspaceCompare("a#c", "b"))
}
