func longestOnes(nums []int, k int) int {
	maxWindow := 0
	curSum := 0
	start := 0

	for end := 0; end < len(nums); end++ {
		curSum += nums[end]

		if end-start+1-curSum <= k {
			maxWindow = max(maxWindow, end-start+1)
		} else {
			curSum -= nums[start]
			start++
		}
	}
	return maxWindow
}
