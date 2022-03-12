package longestPalindrome

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/

/*
@Author: Mickey
*/
func longestPalindrome(s string) string {
	if s == "" || len(s) == 1 {
		return s
	}
	var start, end int

	for i := range s {
		l1, r1 := getMaxLength(s, i, i)
		l2, r2 := getMaxLength(s, i, i+1)
		var l, r int
		maxLen := r1 - l1
		l = l1
		r = r1

		if maxLen < r2-l2 {
			maxLen = r2 - l2
			l = l2
			r = r2
		}
		if maxLen > end-start {
			start = l
			end = r
		}
	}

	return s[start:end]
}

func getMaxLength(s string, left int, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right
}
