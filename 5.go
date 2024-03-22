func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	res := ""

	for left := 0; left < len(s); left++ {
		for right := left + 1; right < len(s)+1; right++ {
			q := s[left:right]
			if len(q) > len(res) {
				if isPalindrome(q) {
					res = s[left:right]
				}
			}
		}
	}
	return res
}
func isPalindrome(s string) bool {
	r := len(s) - 1
	for i := 0; i < len(s); i++ {
		if s[i] != s[r] {
			return false
		}
		r--
	}
	return true
}