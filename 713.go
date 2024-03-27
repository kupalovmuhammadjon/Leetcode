func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	count := 0
	product := 1
	left := 0

	for right, num := range nums {
		product *= num

		// Shrink the window from the left until the product is less than k
		for product >= k {
			product /= nums[left]
			left++
		}

		// The number of subarrays ending at right index is right - left + 1
		count += right - left + 1
	}

	return count
}
