package isPowerOfFour

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfFour(t *testing.T) {
	assert.True(t, isPowerOfFour(16))
	assert.False(t, isPowerOfFour(5))
}
