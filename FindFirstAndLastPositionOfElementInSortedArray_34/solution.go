package FindFirstAndLastPositionOfElementInSortedArray_34

const (
	searchStart = 1
	searchEnd   = 2
)

func searchRange(nums []int, target int) []int {
	return []int{binSearch(nums, target, searchStart), binSearch(nums, target, searchEnd)}
}

func binSearch(nums []int, target int, search byte) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		m := (lo + hi) / 2
		if nums[m] == target {
			if search == searchStart {
				if m == 0 || nums[m-1] < nums[m] {
					return m
				} else {
					hi = m-1
				}
			} else {
				if m == len(nums)-1 || nums[m+1] > nums[m] {
					return m
				} else {
					lo = m+1
				}
			}
		} else if nums[m] < target {
			lo = m+1
		} else {
			hi = m-1
		}
	}
	return -1
}
