package findMedianSortedArrays

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

思路：v1.生成一个新的数组，然后根据奇偶取中间值
2. 二分法，递归
*/

/*
@Author: Mickey
*/

func findMedianSortedArraysV1(nums1 []int, nums2 []int) float64 {
	var num []int
	i := 0
	j := 0
	for i < len(nums1) || j < len(nums2) {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				num = append(num, nums1[i])
				i++
			} else {
				num = append(num, nums2[j])
				j++
			}
			continue
		}
		if i < len(nums1) {
			num = append(num, nums1[i])
			i++
		}
		if j < len(nums2) {
			num = append(num, nums2[j])
			j++
		}
	}

	if len(num) == 0 {
		return 0
	}

	if len(num) == 1 {
		return float64(num[0])
	}

	if len(num)%2 == 0 {
		return float64((num[len(num)/2-1] + num[len(num)/2])) / 2
	} else {
		return float64(num[len(num)/2])
	}

}

const MAX_INT = int(^uint(0) >> 1)

/*
@Author: Mickey
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)

	if n == 0 && m == 0 {
		return 0
	}

	if n == 0 {
		if m%2 == 0 {
			return float64(nums2[m/2]+nums2[m/2-1]) / 2
		}
		return float64(nums2[m/2])
	}

	if m == 0 {
		if n%2 == 0 {
			return float64(nums1[n/2]+nums1[n/2-1]) / 2
		}
		return float64(nums1[n/2])
	}

	var total int = m + n

	if total%2 == 1 {
		return float64(find_kth(nums1, 0, nums2, 0, total/2+1))
	}

	return float64(find_kth(nums1, 0, nums2, 0, total/2)+find_kth(nums1, 0, nums2, 0, total/2+1)) / 2
}

func find_kth(num1 []int, num1Begin int, num2 []int, num2Begin int, k int) int {

	if num1Begin >= len(num1) {
		return num2[num2Begin+k-1]
	}

	if num2Begin >= len(num2) {
		return num1[num1Begin+k-1]
	}

	if k == 1 {
		if num1[num1Begin] < num2[num2Begin] {
			return num1[num1Begin]
		}
		return num2[num2Begin]
	}

	var num1Mid int = MAX_INT
	var num2Mid int = MAX_INT
	if num1Begin+k/2-1 < len(num1) {
		num1Mid = num1[num1Begin+k/2-1]
	}

	if num2Begin+k/2-1 < len(num2) {
		num2Mid = num2[num2Begin+k/2-1]
	}

	if num1Mid < num2Mid {
		return find_kth(num1, num1Begin+k/2, num2, num2Begin, k-k/2)
	}

	return find_kth(num1, num1Begin, num2, num2Begin+k/2, k-k/2)
}
