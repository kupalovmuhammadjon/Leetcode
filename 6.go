func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	mat := make([][]rune, numRows)
	lamp := true
	r := 0
	for i := 0; i < len(s); i++ {
		mat[r] = append(mat[r], rune(s[i]))
		if r == numRows-1 {
			lamp = false
		} else if r == 0 {
			lamp = true
		}
		if lamp {
			r++
		} else {
			r--
		}
	}
	res := ""
	for i := 0; i < len(mat); i++ {
		res += string(mat[i])
	}
	return res
}