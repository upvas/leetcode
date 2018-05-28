package MaximumSubarray_53

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	lastMax := 0
	for _, v := range nums {
		add := 0
		if lastMax > 0 {
			add = lastMax
		}
		lastMax = v + add

		if lastMax > max {
			max = lastMax
		}
	}
	return max
}
