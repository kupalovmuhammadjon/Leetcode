func search(nums []int, target int) int {
	start := 0
	finish := len(nums) - 1

	for start <= finish {
		mid := (start + finish) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			finish = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}