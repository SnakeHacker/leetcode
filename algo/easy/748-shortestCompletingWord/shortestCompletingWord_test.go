package shortestCompletingWord

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestCompletingWord(t *testing.T) {
	licensePlate := "1s3 PSt"
	words := []string{"step", "steps", "stripe", "stepple"}
	actual := shortestCompletingWord(licensePlate, words)
	expected := "steps"

	assert.Equal(t, expected, actual)
}
