package findMedianSortedArrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	expect := 2.5
	actual := findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, expect, actual)

	actual = findMedianSortedArraysV1(nums1, nums2)
	assert.Equal(t, expect, actual)

}

func TestFindMedianSortedArrays2(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	expect := 2.0
	actual := findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, expect, actual)

	actual = findMedianSortedArraysV1(nums1, nums2)
	assert.Equal(t, expect, actual)
}

func TestFindMedianSortedArrays3(t *testing.T) {
	nums1 := []int{0}
	nums2 := []int{}

	expect := 0.0
	actual := findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, expect, actual)

	actual = findMedianSortedArraysV1(nums1, nums2)
	assert.Equal(t, expect, actual)
}

func TestFindMedianSortedArrays4(t *testing.T) {
	nums1 := []int{}
	nums2 := []int{1}

	expect := 1.0
	actual := findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, expect, actual)

	actual = findMedianSortedArraysV1(nums1, nums2)
	assert.Equal(t, expect, actual)
}

func TestFindMedianSortedArrays5(t *testing.T) {
	nums1 := []int{}
	nums2 := []int{}

	expect := 0.0
	actual := findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, expect, actual)

	actual = findMedianSortedArraysV1(nums1, nums2)
	assert.Equal(t, expect, actual)
}

func TestFindMedianSortedArrays6(t *testing.T) {
	nums1 := []int{}
	nums2 := []int{2, 3}

	expect := 2.5
	actual := findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, expect, actual)

	actual = findMedianSortedArraysV1(nums1, nums2)
	assert.Equal(t, expect, actual)
}

func TestFindMedianSortedArrays7(t *testing.T) {
	nums1 := []int{1}
	nums2 := []int{2, 3, 4, 5, 6}

	expect := 3.5
	actual := findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, expect, actual)

	actual = findMedianSortedArraysV1(nums1, nums2)
	assert.Equal(t, expect, actual)
}
