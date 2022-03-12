package getRow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRow(t *testing.T) {
	actual := getRow(2)
	expected := []int{1, 2, 1}
	assert.Equal(t, expected, actual)
}
