package isPalindrome

import (
	"strconv"
)

const MAX_INT = 1<<31 - 1
const MIN_INT = -MAX_INT - 1

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
*/

/*
@Author: Zhang yang
*/

func isPalindromeV1(x int) bool {
	s := strconv.Itoa(x)
	length := len(s)

	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}

	return true
}

/*
@Author: Mickey
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	r := reverse(x)

	return x == r

}

func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}
	if y > MAX_INT || y < MIN_INT {
		return 0
	}

	return y
}
