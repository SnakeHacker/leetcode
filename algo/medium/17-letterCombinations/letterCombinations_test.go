package letterCombinations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	assert.Equal(t, expected, letterCombinations("23"))
}

func TestLetterCombinations2(t *testing.T) {
	expected := []string{}
	assert.Equal(t, expected, letterCombinations(""))
}
