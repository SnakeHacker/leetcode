package array

func Init2DIntArray(nums []int, rows int, columns int) [][]int {
	result := [][]int{}
	r := 0
	i := 0

	for r < rows {
		rs := []int{}
		c := 0
		for c < columns {
			rs = append(rs, nums[i])
			c++
			i++
			if i == len(nums) {
				result = append(result, rs)
				return result
			}
		}
		result = append(result, rs)
		r++
	}
	return result
}
