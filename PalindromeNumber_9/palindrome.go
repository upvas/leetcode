package PalindromeNumber_9

func isPalindrome(x int) bool {
	if 0 <= x && x <= 9 {
		return true
	}

	if x % 10 == 0 {
		return false
	}

	reverse := 0
	for x > reverse {
		reverse = reverse * 10 + x % 10
		x /= 10
	}

	// two possible cases: 1221 and 12921
	return reverse == x || reverse / 10 == x
}