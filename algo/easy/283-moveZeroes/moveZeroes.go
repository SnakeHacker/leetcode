package moveZeroes

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
*/

/*
@Author: Mickey
*/

func moveZeroes2(nums []int) []int {
	l := 0
	for i := range nums {
		if nums[i] != 0 {
			tmp := nums[l]
			nums[l] = nums[i]
			nums[i] = tmp
			l++
		}
	}

	return nums
}

func moveZeroes(nums []int) []int {
	p := 0
	for i, num := range nums {
		if num != 0 {
			if p != i {
				nums[i], nums[p] = nums[p], nums[i]
			}
			p++
		}
	}
	return nums
}
