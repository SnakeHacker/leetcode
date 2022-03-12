package reverseVowels

import (
	"bytes"
)

/*
编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

示例 1:

输入: "hello"
输出: "holle"
示例 2:

输入: "leetcode"
输出: "leotcede"
说明:
元音字母不包含字母"y"。
*/

/*
@Author: Mickey
*/

func reverseVowels(s string) string {
	var vowelIndex []int
	vowelMap := make(map[string]bool)
	vowelMap["a"] = true
	vowelMap["e"] = true
	vowelMap["i"] = true
	vowelMap["o"] = true
	vowelMap["u"] = true
	vowelMap["A"] = true
	vowelMap["E"] = true
	vowelMap["I"] = true
	vowelMap["O"] = true
	vowelMap["U"] = true

	sArray := []rune{}
	for i, ss := range s {
		if _, ok := vowelMap[string(ss)]; ok {
			vowelIndex = append(vowelIndex, i)
		}
		sArray = append(sArray, ss)
	}

	for i := range vowelIndex {
		if i > len(vowelIndex)/2-1 {
			break
		}

		sArray[vowelIndex[i]], sArray[vowelIndex[len(vowelIndex)-1-i]] = sArray[vowelIndex[len(vowelIndex)-1-i]], sArray[vowelIndex[i]]
	}

	var buf bytes.Buffer
	for _, s := range sArray {
		buf.WriteRune(s)
	}

	return buf.String()
}
