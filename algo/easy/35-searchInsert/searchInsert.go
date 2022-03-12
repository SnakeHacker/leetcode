package searchInsert

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0
*/

/*
@Author: Mickey
*/

func searchInsert(nums []int, target int) int {
	if len(nums) == 1 {
		if target < nums[0] {
			return 0
		} else if target == nums[0] {
			return 0
		} else {
			return 1
		}
	}
	mid := nums[len(nums)/2]
	if target == mid {
		return len(nums) / 2
	} else if target > mid {
		return len(nums)/2 + searchInsert(nums[len(nums)/2:], target)
	} else {
		return searchInsert(nums[:len(nums)/2], target)
	}
}
