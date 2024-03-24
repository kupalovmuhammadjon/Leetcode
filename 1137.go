func tribonacci(n int) int {
	f1 := 0
	f2 := 1
	f3 := 1
	for i := 0; i < n; i++ {
		tribo := f1 + f2 + f3
		f1 = f2
		f2 = f3
		f3 = tribo
	}
	return f1
}