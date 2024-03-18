func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	count := 0
	start := 0
	finish := len(nums) - 1

	for start < finish {
		if nums[start]+nums[finish] == k {
			count++
			start++
			finish--
		} else if nums[start]+nums[finish] > k {
			finish--
		} else {
			start++
		}
	}

	return count
}
