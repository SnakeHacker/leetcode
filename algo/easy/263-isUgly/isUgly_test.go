package isUgly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUgly(t *testing.T) {
	assert.True(t, isUgly(1))
	assert.True(t, isUgly(6))
	assert.True(t, isUgly(8))
	assert.False(t, isUgly(14))
	assert.False(t, isUgly(0))
}
