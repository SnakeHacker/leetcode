package sortedSquares

import (
	"sort"
)

/*
给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。

示例 1：

输入：[-4,-1,0,3,10]
输出：[0,1,9,16,100]
示例 2：

输入：[-7,-3,2,3,11]
输出：[4,9,9,49,121]


提示：

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A 已按非递减顺序排序。
*/

/*
@Author: Mickey
*/

func sortedSquares(A []int) []int {
	for i := range A {
		A[i] = A[i] * A[i]
	}

	sort.Ints(A)
	return A
}

func sortedSquares2(A []int) []int {
	neg := []int{}
	pos := []int{}
	for i := range A {
		v := A[i] * A[i]
		if A[i] < 0 {
			neg = append(neg, v)
		} else {
			pos = append(pos, v)
		}
	}

	negIdx := len(neg) - 1
	posIdx := 0
	res := []int{}

	for {
		if negIdx < 0 || posIdx > len(pos)-1 {
			break
		}
		if neg[negIdx] <= pos[posIdx] {
			res = append(res, neg[negIdx])
			negIdx--
		} else {
			res = append(res, pos[posIdx])
			posIdx++
		}
	}

	for negIdx >= 0 {
		res = append(res, neg[negIdx])
		negIdx--
	}

	for posIdx < len(pos) {
		res = append(res, pos[posIdx])
		posIdx++
	}

	return res
}
