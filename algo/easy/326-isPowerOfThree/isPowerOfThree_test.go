package isPowerOfThree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfThree(t *testing.T) {
	assert.True(t, isPowerOfThree(27))
	assert.False(t, isPowerOfThree(6))
}
