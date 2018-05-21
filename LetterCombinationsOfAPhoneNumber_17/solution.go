package LetterCombinationsOfAPhoneNumber_17

var digitsToLetters = [][]string{
	{}, // 0
	{}, // 1
	{"a", "b", "c"},
	{"d", "e", "f"},
	{"g", "h", "i"},
	{"j", "k", "l"},
	{"m", "n", "o"},
	{"p", "q", "r", "s"},
	{"t", "u", "v"},
	{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var res []string

	if len(digits) == 0 {
		return res
	}

	firstDig := digits[0] - 48 // 48 - the code point of "0"
	letters := digitsToLetters[firstDig]

	if len(digits) == 1 {
		return letters
	}

	for _, letter := range letters {
		for _, comb := range letterCombinations(string(digits[1:])) {
			res = append(res, letter + comb)
		}
	}

	return res
}