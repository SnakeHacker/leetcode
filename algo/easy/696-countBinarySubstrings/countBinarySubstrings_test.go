package countBinarySubstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBinarySubstrings(t *testing.T) {
	actual := countBinarySubstrings("10101")
	expected := 4
	assert.Equal(t, expected, actual)
}

func TestCountBinarySubstrings2(t *testing.T) {
	actual := countBinarySubstrings("00110011")
	expected := 6
	assert.Equal(t, expected, actual)
}

func TestCountBinarySubstrings3(t *testing.T) {
	actual := countBinarySubstrings("00")
	expected := 0
	assert.Equal(t, expected, actual)
}

func TestCountBinarySubstrings4(t *testing.T) {
	actual := countBinarySubstrings("100111001")
	expected := 6
	assert.Equal(t, expected, actual)
}
