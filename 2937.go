func findMinimumOperations(s1 string, s2 string, s3 string) int {
	if s1 == s2 && s2 == s3 {
		return 0
	}
	minlen := min(min(len(s1), len(s2)), len(s3))

	for i := minlen; i > 0; i-- {
		if s1[:i] == s2[:i] && s2[:i] == s3[:i] {
			return len(s1) + len(s2) + len(s3) - 3*i
		}
	}

	return -1
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}