package reverseString

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	actual := reverseString([]byte("hello"))
	expected := []byte("olleh")
	assert.Equal(t, expected, actual)
}
