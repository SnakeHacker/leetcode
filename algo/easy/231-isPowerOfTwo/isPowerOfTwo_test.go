package isPowerOfTwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfTwo(t *testing.T) {
	assert.True(t, isPowerOfTwo(1))
	assert.True(t, isPowerOfTwo(16))
	assert.False(t, isPowerOfTwo(218))
}
