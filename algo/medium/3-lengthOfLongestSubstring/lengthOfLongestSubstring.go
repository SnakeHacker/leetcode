package lengthOfLongestSubstring

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

/*
@Author: Mickey
*/

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	var start, end int
	res := start - end + 1

	m := make(map[byte]int)

	for i := range s {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = i
			end = i
		} else {
			cur := end - start + 1
			if cur > res {
				res = cur
			}
			for j := start; j < m[s[i]]; j++ {
				delete(m, s[j])
			}
			start = m[s[i]] + 1
			end = i
			m[s[i]] = i
		}
	}

	cur := end - start + 1
	if cur > res {
		res = cur
	}

	return res
}

func lengthOfLongestSubstringV2(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	start := 0
	end := 0
	res := end - start + 1

	m := make(map[byte]int)
	for i := range s {
		_, ok := m[s[i]]
		end = i

		if ok {
			for j := start; j < end; j++ {
				if s[j] == s[i] {
					start = j + 1
					break
				}
			}
		}

		m[s[i]] = i

		if res < end-start+1 {
			res = end - start + 1
		}
	}

	return res
}
