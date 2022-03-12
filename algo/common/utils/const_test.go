package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConst(t *testing.T) {
	assert.Equal(t, MAX_INT, 2147483647)
	assert.Equal(t, MIN_INT, -2147483648)
}
