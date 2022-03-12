package thirdMax

import (
	. "github.com/SnakeHacker/leetcode/algo/common/utils"
)

/*
给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。

示例 1:

输入: [3, 2, 1]

输出: 1

解释: 第三大的数是 1.
示例 2:

输入: [1, 2]

输出: 2

解释: 第三大的数不存在, 所以返回最大的数 2 .
示例 3:

输入: [2, 2, 3, 1]

输出: 1

解释: 注意，要求返回第三大的数，是指第三大且唯一出现的数。
存在两个值为2的数，它们都排第二。
*/

/*
@Author: Mickey
*/

const minInt = MIN_INT - 1

func thirdMax(nums []int) int {
	max3Flag := false
	max1, max2, max3 := minInt, minInt, minInt
	for _, num := range nums {
		if num >= max1 {
			if num > max1 {
				if max2 > minInt {
					max3Flag = true
				}
				max3 = max2
				max2 = max1
			}
			max1 = num
			continue
		}
		if num >= max2 {
			if num > max2 {
				if max2 > minInt {
					max3Flag = true
				}
				max3 = max2
			}
			max2 = num
			continue
		}
		if num >= max3 {
			if num > minInt {
				max3Flag = true
			}
			max3 = num
			continue
		}
	}
	if max3Flag {
		return max3
	}

	return max1
}
