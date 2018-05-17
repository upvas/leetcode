package CountAndSay_38

import "strconv"

func countAndSay(n int) string {
	if n < 1 {
		return ""
	}

	say := "1"
	for i := 0; i < n-1; i++ {
		cnt := 0
		nextSay := ""
		for j := 0; j < len(say); j++ {
			cnt++
			if j == len(say)-1 || say[j] != say[j+1] {
				nextSay += strconv.Itoa(cnt) + string(say[j])
				cnt = 0
			}
		}
		say = nextSay
	}
	return say
}
