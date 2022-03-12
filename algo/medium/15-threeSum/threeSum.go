package threeSum

import (
	"sort"
)

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

/*
@Author: Mickey
*/

func threeSum(nums []int) [][]int {
	results := [][]int{}

	sort.Ints(nums)
	if len(nums) < 3 {
		return results
	}

	for i := 1; i < len(nums); i++ {
		leftIdx := 0
		rightIdx := len(nums) - 1

		for {
			if leftIdx == i || rightIdx == i {
				break
			}

			left := nums[leftIdx]
			right := nums[rightIdx]

			curr := nums[i]
			sum := curr + left + right
			if sum == 0 {
				results = append(results, []int{left, curr, right})
				rightIdx--
				leftIdx++
			} else if sum < 0 {
				leftIdx++
			} else if sum > 0 {
				rightIdx--
			}
		}
	}

	uniqueResults := [][]int{}

	for _, r := range results {
		flag := false
		for _, ur := range uniqueResults {
			if r[0] == ur[0] && r[1] == ur[1] && r[2] == ur[2] {
				flag = true
				break
			}
		}
		if !flag {
			uniqueResults = append(uniqueResults, r)
		}
	}

	return uniqueResults
}
