func maxSubarrayLength(nums []int, k int) int {
	maxLen := 0
	l := 0
	freq := map[int]int{}
	for r := 0; r < len(nums); r++ {
		freq[nums[r]]++
		for freq[nums[r]] > k {
			freq[nums[l]]--
			l++
		}

		if maxLen < r-l+1 {
			maxLen = r - l + 1
		}

	}
	return maxLen
}
