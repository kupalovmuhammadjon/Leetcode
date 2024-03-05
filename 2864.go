func maximumOddBinaryNumber(s string) string {
	zeros := ""
	ones := ""
	for _, c := range s {
		if c == '1' {
			ones += "1"
		} else {
			zeros += "0"
		}
	}
	if len(ones) == 1 {
		return zeros + ones
	}

	return ones[1:] + zeros + "1"

}