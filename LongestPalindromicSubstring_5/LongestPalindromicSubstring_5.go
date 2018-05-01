package LongestPalindromicSubstring_5

// Brute Force Solution - O(n^3)
func longestPalindrome(s string) string {
	r := []rune(s)
	var lp []rune
	for i := 0; i < len(r); i++ {
		for j := i; j < len(r); j++ {
			if isPalindrome(r[i:j+1]) && len(lp) < (j-i+1) {
				lp = r[i:j+1]
			}
		}
	}
	return string(lp)
}

func isPalindrome(s []rune) bool {
	for i := 0; i <= (len(s) - 1)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}