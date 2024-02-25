func countPrimeSetBits(left int, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		bits := 0
		n := i
		for n > 0 {
			if n%2 == 1 {
				bits++
			}
			n /= 2
		}
		if isPrime(bits) {
			count++
		}
	}
	return count
}
func isPrime(n int) bool {
	if n == 1 {
		return false
	} else if n == 2 {
		return true
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}