package RemoveDuplicatesFromSortedArray_26

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	p := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		p++
		nums[p] = nums[i]
	}
	return p+1
}