package reverseStr

import "bytes"

/*
给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。如果剩余少于 k 个字符，则将剩余的所有全部反转。如果有小于 2k 但大于或等于 k 个字符，则反转前 k 个字符，并将剩余的字符保持原样。

示例:

输入: s = "abcdefg", k = 2
输出: "bacdfeg"
要求:

该字符串只包含小写的英文字母。
给定字符串的长度和 k 在[1, 10000]范围内。
*/

/*
@Author: Mickey
*/

func reverseStr(s string, k int) string {
	if k <= 0 {
		return s
	}
	var buf bytes.Buffer
	tmp := []rune{}

	if len(s) < k {
		for j := range s {
			buf.WriteRune(rune(s[len(s)-1-j]))
		}
		return buf.String()
	}

	for i, c := range s {
		if i%(2*k) < k {
			tmp = append(tmp, c)
			continue
		} else {
			if len(tmp) > 0 {
				for j := range tmp {
					buf.WriteRune(tmp[len(tmp)-1-j])
				}
				tmp = []rune{}
			}
			buf.WriteRune(c)
		}
	}

	if len(tmp) > 0 {
		for j := range tmp {
			buf.WriteRune(tmp[len(tmp)-1-j])
		}
	}

	return buf.String()
}
