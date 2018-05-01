package _Sum_15

import (
	"sort"
	"fmt"
)

// [-1, -1, -1, -1, 0, 0, 0, 0, 0, 0, 2, 2, 2, 2, 2],
// -1, -1, 0, 0, 0, 2, 2
// c = -a - b
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	fmt.Println(nums)

	zeroCnt := 0

	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		if a == 0 {
			zeroCnt++
			if zeroCnt > 2 {
				res = append(res, []int{0, 0, 0})
				break
			}
			continue
		}

		if nums[i+1] == a && nums[i+2] == a {
			continue
		}

		seen := make(map[int]bool, len(nums)-i)
		for j := i+1; j < len(nums); j++ {
			c := nums[j]
			if c == a || j < len(nums)-1 && nums[j+1] == c {
				continue
			}

			b := -1*(a+c)

			if seen[b] {
				res = append(res, []int{a, b, c})
			} else {
				seen[c] = true
			}
		}
	}

	return res
}