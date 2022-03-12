package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	actual := MaxInt(1, 2)
	expected := 2
	assert.Equal(t, expected, actual)
}

func TestMax2(t *testing.T) {
	actual := MaxInt(2, 1)
	expected := 2
	assert.Equal(t, expected, actual)
}

func TestMin(t *testing.T) {
	actual := MinInt(1, 2)
	expected := 1
	assert.Equal(t, expected, actual)
}

func TestMin2(t *testing.T) {
	actual := MinInt(2, 1)
	expected := 1
	assert.Equal(t, expected, actual)
}
