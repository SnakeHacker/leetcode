package maxArea

/*
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例:

输入: [1,8,6,2,5,4,8,3,7]
输出: 49
*/

/*
@Author: Mickey
*/

func maxArea2(height []int) int {
	maxArea := 0
	if len(height) < 2 {
		return maxArea
	}

	left := 0
	right := len(height) - 1

	for {
		currentHeight := height[left]
		if height[right] < currentHeight {
			currentHeight = height[right]
		}

		if right == left {
			break
		}

		if (right-left)*currentHeight > maxArea {
			maxArea = (right - left) * currentHeight
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	left := 0
	right := len(height) - 1

	maxSum := 0

	if height[left] < height[right] {
		maxSum = (right - left) * height[left]
	} else {
		maxSum = (right - left) * height[right]
	}

	for {
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

		if left >= right {
			break
		}

		sum := 0

		if height[left] < height[right] {
			sum = (right - left) * height[left]
		} else {
			sum = (right - left) * height[right]
		}

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum

}
