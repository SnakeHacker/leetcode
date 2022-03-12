package defangIPaddr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefangIPaddr(t *testing.T) {
	actual := defangIPaddr("1.1.1.1")
	expected := "1[.]1[.]1[.]1"
	assert.Equal(t, expected, actual)
}
