package ImplementStrStr_28


func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	OUTER:
	for i := 0; i <= len(haystack) - len(needle); i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				continue OUTER
			}
		}
		return i
	}

	return -1
}