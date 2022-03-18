package quickSort

func quickSort(arr *[]int, left int, right int) []int {
	if left >= right {
		return *arr
	}
	privot := partition(arr, left, right)
	quickSort(arr, left, privot-1)
	quickSort(arr, privot+1, right)
	return *arr
}

// 2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9
func partition(arr *[]int, left int, right int) int {
	privot := (*arr)[right]

	i := left - 1
	for j := left; j < right; j++ {
		if (*arr)[j] < privot {
			i++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}

	(*arr)[i+1], (*arr)[right] = (*arr)[right], (*arr)[i+1]
	return i + 1
}
