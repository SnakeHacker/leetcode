package reverse

import (
	"strconv"

	"github.com/SnakeHacker/leetcode/algo/common/utils"
)

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/

/*
@Author: Zhangyang
*/

func reverseV1(x int) int {
	s := strconv.Itoa(x)
	if s[0] == '-' {
		v, _ := strconv.ParseInt(reverseString(s[1:]), 10, 64)

		if v-1 < utils.MAX_INT {
			return int(-v)
		} else {
			return 0
		}
	} else {
		v, _ := strconv.ParseInt(reverseString(s), 10, 64)

		if v < utils.MAX_INT {
			return int(v)
		} else {
			return 0
		}
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

/*
@Author: Mickey
*/

func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}
	if y > utils.MAX_INT || y < utils.MIN_INT {
		return 0
	}

	return y
}
