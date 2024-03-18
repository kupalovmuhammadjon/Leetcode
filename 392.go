func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	} else if len(t) == 0 {
		return false
	}
	k := 0
	covered := 0
	for i := 0; i < len(t); i++ {
		if s[k] == t[i] {
			covered++
			k++
			if covered == len(s) {
				return true
			}
		}
	}
	return covered == len(s)
}