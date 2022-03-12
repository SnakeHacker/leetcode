package hammingDistance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingDistance(t *testing.T) {
	actual := hammingDistance(1, 4)
	expeted := 2
	assert.Equal(t, expeted, actual)
}
