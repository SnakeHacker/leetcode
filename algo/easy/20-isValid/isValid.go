package isValid

import . "github.com/SnakeHacker/leetcode/algo/common/stack"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/

/*
@Author: Mickey
*/

func isValid(s string) bool {
	/*
		ASCII
		{ 123	} 125
		[ 91	] 93
		( 40	) 41
	*/
	stack := InitStack()
	for _, c := range s {
		if c == 123 || c == 91 || c == 40 {
			stack.Push(int(c))
		}
		if c == 125 || c == 93 || c == 41 {
			p := stack.Pop()
			if p == nil {
				return false
			}
			if !((c == 125 && p == 123) || (c == 93 && p == 91) || (c == 41 && p == 40)) {
				return false
			}
		}
	}

	if stack.Pop() == nil {
		return true
	}

	return false
}
