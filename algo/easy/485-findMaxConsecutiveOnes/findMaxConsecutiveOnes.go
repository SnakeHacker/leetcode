package findMaxConsecutiveOnes

/*
给定一个二进制数组， 计算其中最大连续1的个数。

示例 1:

输入: [1,1,0,1,1,1]
输出: 3
解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
注意：

输入的数组只包含 0 和1。
输入数组的长度是正整数，且不超过 10,000。
*/

/*
@Author: Mickey
*/

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	now := 0
	for _, num := range nums {
		if num == 1 {
			now++
			if now > max {
				max = now
			}
		} else {
			now = 0
		}
	}
	return max
}
