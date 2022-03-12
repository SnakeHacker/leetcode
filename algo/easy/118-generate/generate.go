package generate

/*
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

/*
@Author: Mickey
*/

func generate(numRows int) [][]int {
	result := [][]int{}

	line1 := []int{1}
	line2 := []int{1, 1}

	if numRows == 0 {
		return result
	}

	if numRows >= 1 {
		result = append(result, line1)
	}

	if numRows >= 2 {
		result = append(result, line2)
	}

	if numRows >= 3 {
		for i := 2; i < numRows; i++ {
			result = append(result, generateNewLine(result[i-1]))
		}
	}

	return result
}

func generateNewLine(oldLine []int) (newLine []int) {
	oldLen := len(oldLine)

	newLine = append(newLine, 1)
	for i := 1; i < oldLen; i++ {
		newLine = append(newLine, oldLine[i]+oldLine[i-1])
	}
	newLine = append(newLine, 1)

	return newLine
}
