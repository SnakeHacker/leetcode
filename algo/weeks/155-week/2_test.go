package week

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDietPlanPerformance(t *testing.T) {
	actual := dietPlanPerformance([]int{1, 2, 3, 4, 5}, 1, 3, 3)
	expected := 0

	assert.Equal(t, expected, actual)
}

func TestDietPlanPerformanc2(t *testing.T) {
	actual := dietPlanPerformance([]int{3, 2}, 2, 0, 1)
	expected := 1

	assert.Equal(t, expected, actual)
}

func TestDietPlanPerformanc3(t *testing.T) {
	actual := dietPlanPerformance([]int{6, 5, 0, 0}, 2, 1, 5)
	expected := 0

	assert.Equal(t, expected, actual)
}

func TestDietPlanPerformanc4(t *testing.T) {
	actual := dietPlanPerformance([]int{6, 13, 8, 7, 10, 1, 12, 11}, 6, 5, 37)
	expected := 3

	assert.Equal(t, expected, actual)
}
