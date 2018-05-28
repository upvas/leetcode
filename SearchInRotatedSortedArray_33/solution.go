package SearchInRotatedSortedArray_33

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	p := getPivot(nums, 0, len(nums)-1)
	if r := binarySearch(nums, p, len(nums)-1, target); r != -1 {
		return r
	}

	return binarySearch(nums, 0, p-1, target)
}

func getPivot(nums []int, s, e int) int {
	if s == e {
		return s
	}

	m := (s+e)/2

	if s < m && m < e && nums[m-1] > nums[m] && nums[m+1] >= nums[m] {
		return m
	} else if nums[s] > nums[m] {
		return getPivot(nums, s, m-1)
	} else if nums[e] < nums[m] {
		return getPivot(nums, m+1, e)
	} else {
		return s
	}
}

func binarySearch(nums []int, s, e int, target int) int {
	if e < s {
		return -1
	}

	m := (s+e)/2
	if nums[m] == target {
		return m
	} else if nums[m] > target {
		return binarySearch(nums, s, m-1, target)
	} else {
		return binarySearch(nums, m+1, e, target)
	}
}