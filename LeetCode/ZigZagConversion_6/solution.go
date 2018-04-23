package ZigZagConversion_6

func convert(s string, numRows int) string {
	runes := []rune(s)
	if len(runes) <= numRows || numRows == 1 {
		return string(runes)
	}

	result := make([]rune, 0, len(runes))
	k := 2*numRows-2

	for i := 0; i < numRows; i++ {
		for j := 0; ; j++ {
			idx := k*j+i
			if idx >= len(runes) {
				break
			}
			result = append(result, runes[idx])

			if i == 0 || i == numRows-1 {
				continue
			}

			idx = k*(j+1)-i
			if idx >= len(runes) {
				break
			}
			result = append(result, runes[idx])
		}
	}

	return string(result)
}