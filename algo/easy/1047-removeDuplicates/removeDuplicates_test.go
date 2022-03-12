package removeDuplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	actual := removeDuplicates("abbaac")
	expected := "ac"
	assert.Equal(t, expected, actual)
}
