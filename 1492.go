func kthFactor(n int, k int) int {
	cur := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			cur++
			if cur == k {
				return i
			}
		}
	}
	return -1
}