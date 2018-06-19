package GenerateParentheses_22

func generateParenthesis(n int) []string {
	return gen("", n, 0)
}

func gen(ch string, n int, b int) []string {
	if n == 0 || b < 0 {
		return []string{}
	}

	var res []string

	if n < b {
		for _, comb := range gen("(", n-1, b+1) {
			if len(comb) == n-1 {
				res = append(res, ch + comb)
			}
		}
	}

	if b > 0 {
		for _, comb := range gen(")", n-1, b-1) {
			if len(comb) == n-1 {
				res = append(res, ch + comb)
			}
		}
	}

	return res
}