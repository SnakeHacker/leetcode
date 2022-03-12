package addStrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddStrings(t *testing.T) {
	actual := addStrings("123", "456")
	expected := "579"
	assert.Equal(t, expected, actual)
}

func TestAddStrings2(t *testing.T) {
	actual := addStrings("999", "1")
	expected := "1000"
	assert.Equal(t, expected, actual)
}
