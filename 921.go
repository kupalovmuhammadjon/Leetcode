func minAddToMakeValid(s string) int {
	opener := 0
	closer := 0
	for _, p := range s {
		if p == '(' {
			opener++
		} else {
			if opener > 0 {
				opener--
			} else {
				closer++
			}
		}
	}
	return opener + closer
}
