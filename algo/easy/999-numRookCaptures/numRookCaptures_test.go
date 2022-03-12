package numRookCaptures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumRookCaptures(t *testing.T) {
	board := [][]byte{
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'p', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'p', '.', '.', '.', '.'},
		[]byte{'p', 'p', '.', 'R', '.', 'p', 'B', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'B', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'p', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	actual := numRookCaptures(board)
	expected := 3
	assert.Equal(t, expected, actual)
}

func TestNumRookCaptures2(t *testing.T) {
	board := [][]byte{
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'p', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'p', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	actual := numRookCaptures(board)
	expected := 3
	assert.Equal(t, expected, actual)
}

func TestNumRookCaptures3(t *testing.T) {
	board := [][]byte{
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		[]byte{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		[]byte{'.', 'p', 'B', 'R', 'R', 'p', '.', '.'},
		[]byte{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		[]byte{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	actual := numRookCaptures(board)
	expected := 0
	assert.Equal(t, expected, actual)
}

func TestNumRookCaptures4(t *testing.T) {
	board := [][]byte{
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', 'R', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	actual := numRookCaptures(board)
	expected := 0
	assert.Equal(t, expected, actual)
}
