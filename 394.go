func decodeString(s string) string {
	stack := []interface{}{}
	res := ""
	num := 0
	var letter string
	i := 0
	for i < len(s) {
		if unicode.IsDigit(rune(s[i])) {
			num = int(s[i]) - '0'
			var j int
			for j = i + 1; j < len(s); j++ {
				if unicode.IsDigit(rune(s[j])) {
					num = num*10 + int(s[j]) - '0'
				} else {
					break
				}
			}
			stack = append(stack, num)
			i = j
			num = 0
		} else if unicode.IsLetter(rune(s[i])) {
			letter = string(s[i])
			var j int
			for j = i + 1; j < len(s); j++ {
				if unicode.IsLetter(rune(s[j])) {
					letter += string(s[j])
				} else {
					break
				}
			}
			stack = append(stack, letter)
			i = j
			letter = ""
		} else if s[i] == '[' {
			stack = append(stack, '[')
			i++
		} else if s[i] == ']' {
			st := ""
			for len(stack) > 0 {
				if str, ok := stack[len(stack)-1].(string); ok {
					s := str
					st = s + st
					stack = stack[:len(stack)-1]
				} else if stack[len(stack)-1] == '[' {
					stack = stack[:len(stack)-1]
					break
				}
			}
			put := strings.Repeat(st, stack[len(stack)-1].(int))
			stack = stack[:len(stack)-1]
			stack = append(stack, put)
			i++
		} else {
			i++
		}
	}
	for _, v := range stack {
		res += v.(string)
	}

	return res
}