package IntegerToRoman_12

import "strings"

func intToRoman(num int) string {
	if !(1 <= num && num <= 3999) {
		return ""
	}

	romans := make([]string, 0)
	base := 0
	for ; num > 0; {
		d := num%10
		romansD := getRomanDigs(base, d)
		romans = append(romansD, romans...)
		num /= 10
		base++
	}
	return strings.Join(romans, "")
}


var intToRomanInts = map[int][]int{
	0: {},
	1: {1},
	2: {1, 1},
	3: {1, 1, 1},
	4: {1, 5},
	5: {5},
	6: {5, 1},
	7: {5, 1, 1},
	8: {5, 1, 1, 1},
	9: {1, 10},
}
var romanDigits = map[int]map[int]string{
	0 : {1: "I", 5: "V", 10: "X"},
	1 : {1: "X", 5: "L", 10: "C"},
	2 : {1: "C", 5: "D", 10: "M"},
	3 : {1: "M",},
}
func getRomanDigs(base, dig int) (res []string) {
	for _, romanInt := range intToRomanInts[dig] {
		res = append(res, romanDigits[base][romanInt])
	}
	return
}