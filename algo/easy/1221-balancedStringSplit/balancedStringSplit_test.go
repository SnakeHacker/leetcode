package balancedStringSplit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalancedStringSplit(t *testing.T) {
	actual := balancedStringSplit("RLRRLLRLRL")
	assert.Equal(t, 4, actual)

	actual = balancedStringSplit("RLLLLRRRLR")
	assert.Equal(t, 3, actual)

	actual = balancedStringSplit("LLLLRRRR")
	assert.Equal(t, 1, actual)
}
