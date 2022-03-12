package pivotIndex

/*
给定一个整数类型的数组 nums，请编写一个能够返回数组“中心索引”的方法。

我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。

如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。

示例 1:

输入:
nums = [1, 7, 3, 6, 5, 6]
输出: 3
解释:
索引3 (nums[3] = 6) 的左侧数之和(1 + 7 + 3 = 11)，与右侧数之和(5 + 6 = 11)相等。
同时, 3 也是第一个符合要求的中心索引。
示例 2:

输入:
nums = [1, 2, 3]
输出: -1
解释:
数组中不存在满足此条件的中心索引。
说明:

nums 的长度范围为 [0, 10000]。
任何一个 nums[i] 将会是一个范围在 [-1000, 1000]的整数。
*/

/*
@Author: Mickey
*/

func pivotIndexV1(nums []int) int {
	for i := range nums {
		left := 0
		right := 0
		for j := 0; j < i; j++ {
			left += nums[j]
		}

		for k := i + 1; k < len(nums); k++ {
			right += nums[k]
		}
		if left == right {
			return i
		}
	}
	return -1
}

/*
@Author: Mickey
*/

func pivotIndexV2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	leftMap := make(map[int]int)
	rightMap := make(map[int]int)
	for i := range nums {
		left := 0
		right := 0

		l, ok := leftMap[i-2]
		if ok {
			left = l + nums[i-1]
			leftMap[i-1] = left
		} else {
			if i == 1 {
				leftMap[i-1] = nums[0]
				left = leftMap[i-1]
			}
		}

		r, ok := rightMap[i]
		if ok {
			right = r - nums[i]
			rightMap[i+1] = right
		} else {
			if i < len(nums) {
				sum := 0
				for j := i; j < len(nums); j++ {
					sum += nums[j]
				}
				rightMap[i] = sum
				right = rightMap[i] - nums[i]
			}
		}

		if left == right {
			return i
		}
	}

	return -1
}

/*
@Author: Mickey
*/

func pivotIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	leftMap := make(map[int]int)
	for i := range nums {
		left := 0
		l, ok := leftMap[i-2]
		if ok {
			left = l + nums[i-1]
			leftMap[i-1] = left
		} else {
			if i == 1 {
				leftMap[i-1] = nums[0]
				left = leftMap[i-1]
			}
		}
		if sum-nums[i]-2*left == 0 {
			return i
		}
	}

	return -1
}
