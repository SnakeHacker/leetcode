package peakIndexInMountainArray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	actual := peakIndexInMountainArray([]int{0, 1, 0})
	expected := 1
	assert.Equal(t, expected, actual)
}

func TestPeakIndexInMountainArray2(t *testing.T) {
	actual := peakIndexInMountainArray([]int{0, 2, 1, 0})
	expected := 1
	assert.Equal(t, expected, actual)
}
