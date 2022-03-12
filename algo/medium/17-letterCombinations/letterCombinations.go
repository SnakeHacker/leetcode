package letterCombinations

var digitsMap = map[string][]string{
	"1": []string{},
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

var combinations []string

func letterCombinations(digits string) []string {
	combinations = []string{}
	if len(digits) == 0 {
		return combinations
	}

	backtrack(digits, 0, "")

	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
		return
	}

	digit := string(digits[index])
	letters := digitsMap[digit]

	for i := 0; i < len(letters); i++ {
		backtrack(digits, index+1, combination+string(letters[i]))
	}
}
