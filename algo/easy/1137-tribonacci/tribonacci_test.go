package tribonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTribonacci(t *testing.T) {
	actual := tribonacci(4)
	expected := 4
	assert.Equal(t, expected, actual)
}

func TestTribonacci2(t *testing.T) {
	actual := tribonacci(25)
	expected := 1389537
	assert.Equal(t, expected, actual)
}
