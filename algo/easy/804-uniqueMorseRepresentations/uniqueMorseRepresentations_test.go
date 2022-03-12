package uniqueMorseRepresentations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueMorseRepresentations(t *testing.T) {
	actual := uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"})
	expected := 2
	assert.Equal(t, expected, actual)
}
