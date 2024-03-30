func minSubArrayLen(target int, nums []int) int {
	l := 0
	minLen := len(nums) + 1
	curSum := 0
	for r := 0; r < len(nums); r++ {
		curSum += nums[r]
		for curSum >= target {
			if minLen > r-l+1 {
				minLen = r - l + 1
			}
			curSum -= nums[l]
			l++
		}
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}