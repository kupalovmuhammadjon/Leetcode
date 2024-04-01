func lengthOfLastWord(s string) int {
	cur := 0
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsLetter(rune(s[i])) {
			cur++
		} else {
			if cur > 0 {
				return cur
			}
		}
	}
	if cur > 0 {
		return cur
	}
	return 0
}