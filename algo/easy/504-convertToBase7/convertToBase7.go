package convertToBase7

import (
	"bytes"
	"strconv"
)

/*
给定一个整数，将其转化为7进制，并以字符串形式输出。

示例 1:

输入: 100
输出: "202"
示例 2:

输入: -7
输出: "-10"
注意: 输入范围是 [-1e7, 1e7] 。

*/

/*
@Author: Mickey
*/

func convertToBase7(num int) string {
	var result bytes.Buffer
	if num < 0 {
		result.WriteString("-")
		num = -1 * num
	}
	var res []string
	for {
		res = append(res, strconv.Itoa(num%7))

		if num/7 == 0 {
			break
		}
		num = num / 7
	}

	for i := range res {
		result.WriteString(res[len(res)-1-i])
	}
	return result.String()
}
