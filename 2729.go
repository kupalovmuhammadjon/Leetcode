func isFascinating(n int) bool {
	s := "123456789"
	num := strconv.Itoa(n)
	num += strconv.Itoa(n * 2)
	num += strconv.Itoa(n * 3)
	fmt.Println(num)
	if strings.Contains(num, "0") || len(num) > 9 {
		return false
	}
	for _, v := range s {
		if !strings.ContainsRune(num, v) {
			return false
		}
	}
	return true
}