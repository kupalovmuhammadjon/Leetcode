func maxArea(height []int) int {
	maxAmount := -123333
	i := 0
	j := len(height) - 1

	for i != j && i < j {
		curAmount := min(height[i], height[j]) * (j - i)
		if maxAmount < curAmount {
			maxAmount = curAmount
		} else if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxAmount
}