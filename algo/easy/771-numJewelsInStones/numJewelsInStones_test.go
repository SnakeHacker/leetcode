package numJewelsInStones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumJewelsInStones(t *testing.T) {
	J := "aA"
	S := "aAAbbbb"
	actual := numJewelsInStones(J, S)
	expected := 3
	assert.Equal(t, expected, actual)
}
