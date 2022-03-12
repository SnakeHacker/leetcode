package toLowerCase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToLowerCase(t *testing.T) {
	actual := toLowerCase("hehe")
	expected := "hehe"
	assert.Equal(t, expected, actual)
}

func TestToLowerCase2(t *testing.T) {
	actual := toLowerCase("HeAzZ")
	expected := "heazz"
	assert.Equal(t, expected, actual)
}
