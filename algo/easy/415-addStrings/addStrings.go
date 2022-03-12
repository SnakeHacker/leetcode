package addStrings

import (
	"bytes"
	"strconv"
)

/*
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

*/

/*
@Author: Mickey
*/

func addStrings(num1 string, num2 string) string {
	var buf bytes.Buffer
	l1 := len(num1)
	l2 := len(num2)

	carry := false
	var result []string

	for i, j := l1-1, l2-1; i >= 0 || j >= 0; {
		digit := 0
		if carry {
			digit++
			carry = false
		}

		if i >= 0 {
			digit += int(num1[i]) - 48
			i--
		}
		if j >= 0 {
			digit += int(num2[j]) - 48
			j--
		}

		if digit >= 10 {
			carry = true
		}
		result = append(result, strconv.Itoa(digit%10))
	}

	if carry {
		result = append(result, "1")
	}

	for i := range result {
		buf.WriteString(result[len(result)-1-i])
	}

	return buf.String()
}
