package week

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	var result int
	for i := range calories {
		tmp := 0
		out := false
		for j := 0; j < k; j++ {
			if i+j > len(calories)-1 {
				out = true
				break
			}
			tmp += calories[i+j]
		}
		if out {
			break
		}

		if tmp > upper {
			result++
		}
		if tmp < lower {
			result--
		}
	}

	return result
}
