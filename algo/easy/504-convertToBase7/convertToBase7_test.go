package convertToBase7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToBase7(t *testing.T) {
	actual := convertToBase7(100)
	expected := "202"
	assert.Equal(t, expected, actual)
}

func TestConvertToBase72(t *testing.T) {
	actual := convertToBase7(-7)
	expected := "-10"
	assert.Equal(t, expected, actual)
}
