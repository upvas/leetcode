package SearchInsertPosition_35

func searchInsert(nums []int, target int) int {
	lo := 0
	hi := len(nums)-1
	for lo <= hi {
		m := (lo+hi)/2
		if nums[m] == target {
			var i int
			for i = m-1; i >= 0 && nums[i] == nums[m]; i-- {}
			return i+1
		} else if nums[m] < target {
			lo = m+1
		} else {
			hi = m-1
		}
	}

	return hi+1
}