package ContainerWithMostWater_11

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	res := 0
	l, r := 0, len(height)-1
	for ;l < r; {
		area := max(height[l], height[r]) * (r-l)
		if area > res {
			res = area
		}

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}