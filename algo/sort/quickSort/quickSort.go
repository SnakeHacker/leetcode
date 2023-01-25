package quickSort

func quickSort(arr *[]int, left int, right int) []int {
	if left >= right {
		return *arr
	}

	privot := partition(arr, left, right)
	quickSort(arr, left, privot)
	quickSort(arr, privot+1, right)
	return *arr
}

func partition(arr *[]int, left int, right int) int {
	i := left - 1
	for j := left; j < right; j++ {
		if (*arr)[j] < (*arr)[right] {
			i++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}

	(*arr)[i+1], (*arr)[right] = (*arr)[right], (*arr)[i+1]

	return i
}
