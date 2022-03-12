package longestPalindrome

/*
给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。

在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。

注意:
假设字符串的长度不会超过 1010。

示例 1:

输入:
"abccccdd"

输出:
7

解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
*/

/*
@Author: Mickey
*/
func longestPalindrome(s string) int {
	var result int

	if len(s) < 2 {
		return len(s)
	}

	m := make(map[rune]int)
	for _, c := range s {
		m[c] += 1
	}

	odd := 0
	for _, v := range m {
		switch v % 2 {
		case 1:
			odd = 1
			if v > 1 {
				result += v - 1
			}
		case 0:
			result += v
		}
	}

	return result + odd
}
