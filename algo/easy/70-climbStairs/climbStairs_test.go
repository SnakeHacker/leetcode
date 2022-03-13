package climbStairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	assert.Equal(t, 2, climbStairs(2))
}

func TestClimbStairs2(t *testing.T) {
	assert.Equal(t, 3, climbStairs(3))
}

func TestClimbStairs3(t *testing.T) {
	assert.Equal(t, 5, climbStairs(4))
}
