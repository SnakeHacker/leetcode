package findRestaurant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRestaurant(t *testing.T) {
	l1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	l2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}

	actual := findRestaurant(l1, l2)
	expected := []string{"Shogun"}

	assert.Equal(t, expected, actual)
}

func TestFindRestaurant2(t *testing.T) {
	l1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	l2 := []string{"KFC", "Shogun", "Burger King"}

	actual := findRestaurant(l1, l2)
	expected := []string{"Shogun"}

	assert.Equal(t, expected, actual)
}
