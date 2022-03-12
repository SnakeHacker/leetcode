package maxSubArray

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/

/*
@Author: Mickey
*/

func maxSubArray2(nums []int) int {
	opt := make(map[int]int)
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum := 0
	for i, num := range nums {
		if i == 0 {
			opt[i] = num
			maxSum = opt[i]
		} else {
			a := num
			b := num + opt[i-1]
			max := 0
			if a > b {
				max = a
			} else {
				max = b
			}

			opt[i] = max
			if opt[i] > maxSum {
				maxSum = opt[i]
			}
		}
	}

	return maxSum
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numsM := []int{}
	for i, num := range nums {
		if i == 0 {
			numsM = append(numsM, num)
			continue
		}

		if numsM[i-1] > 0 {
			numsM = append(numsM, numsM[i-1]+num)
		} else {
			numsM = append(numsM, num)
		}
	}

	max := numsM[0]

	for _, num := range numsM {
		if max < num {
			max = num
		}
	}
	return max
}
