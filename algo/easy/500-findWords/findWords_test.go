package findWords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWords(t *testing.T) {
	actual := findWords([]string{"Hello", "Alaska", "Dad", "Peace"})
	expected := []string{"Alaska", "Dad"}
	assert.Equal(t, expected, actual)
}
