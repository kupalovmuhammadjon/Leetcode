func binaryGap(n int) int {
	maxGap := 0
	prevOneIndex := -1
	currentIndex := 0

	for n > 0 {
		current := n % 2

		if current == 1 {
			if prevOneIndex != -1 {
				gap := currentIndex - prevOneIndex
				if gap > maxGap {
					maxGap = gap
				}
			}
			prevOneIndex = currentIndex
		}

		currentIndex++
		n /= 2
	}

	return maxGap
}
