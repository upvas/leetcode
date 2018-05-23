package MaximumSubarray_53

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for i := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}
