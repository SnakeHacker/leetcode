package reverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse0(t *testing.T) {
	expect := 321
	assert.Equal(t, expect, reverseV1(123))
	assert.Equal(t, expect, reverse(123))
}

func TestReverse1(t *testing.T) {
	expect := -321
	assert.Equal(t, expect, reverseV1(-123))
	assert.Equal(t, expect, reverse(-123))
}

func TestReverse2(t *testing.T) {
	expect := 21
	assert.Equal(t, expect, reverseV1(120))
	assert.Equal(t, expect, reverse(120))
}

func TestReverse3(t *testing.T) {
	expect := 0
	assert.Equal(t, expect, reverseV1(1534236469))
	assert.Equal(t, expect, reverse(1534236469))
}
