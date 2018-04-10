package LongestCommonPrefix_14

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || strs[0] == "" {
		return ""
	}

	prefix := []rune(strs[0]) // convert to runes for rune-to-rune comparison below
	for i := 1; i < len(strs) && len(prefix) > 0; i++ { // while common prefix is not empty
		cnt := 0 // count of common runes
		for j, r := range []rune(strs[i]) {
			if j >= len(prefix) || r != prefix[j] {
				break
			}
			cnt++
		}
		prefix = prefix[:cnt]
	}
	return string(prefix)
}