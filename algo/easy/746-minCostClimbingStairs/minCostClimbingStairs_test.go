package minCostClimbingStairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	actual := minCostClimbingStairs([]int{10, 15, 20})
	expected := 15
	assert.Equal(t, expected, actual)
}

func TestMinCostClimbingStairs2(t *testing.T) {
	actual := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	expected := 6
	assert.Equal(t, expected, actual)
}
