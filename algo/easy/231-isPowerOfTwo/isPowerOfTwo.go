package isPowerOfTwo

import "math"

/*
示例 1:

输入: 1
输出: true
解释: 2^0 = 1
示例 2:

输入: 16
输出: true
解释: 2^4 = 16
示例 3:

输入: 218
输出: false
*/

/*
@Author: Mickey
*/

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	// raise speed
	if n%2 != 0 {
		return false
	}
	for i, p := 0, 0; p <= n; i++ {
		p = int(math.Pow(2, float64(i)))
		if p == n {
			return true
		}
	}

	return false
}
