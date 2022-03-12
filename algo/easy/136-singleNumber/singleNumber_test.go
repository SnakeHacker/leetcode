package singleNumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	acutal := singleNumber([]int{2, 2, 1})
	expected := 1
	assert.Equal(t, expected, acutal)
}

func TestSingleNumber2(t *testing.T) {
	acutal := singleNumber([]int{4, 1, 2, 1, 2})
	expected := 4
	assert.Equal(t, expected, acutal)
}
