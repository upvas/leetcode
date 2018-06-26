package PlusOne_66

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	c := 1
	for i := len(digits)-1; i >= 0 && c > 0; i-- {
		digits[i] += c
		c = digits[i] / 10
		digits[i] %= 10
	}

	if c > 0 {
		digits = append([]int{c}, digits...)
	}

	return digits
}