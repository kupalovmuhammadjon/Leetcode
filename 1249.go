func minRemoveToMakeValid(s string) string {
    var result strings.Builder 
    balance := 0
    for _, c := range s {
        if c == '(' {
            balance++
        } else if c == ')' {
            if balance == 0 {
                continue
            }
            balance--
        }
        result.WriteRune(c)
    }
    res := result.String()
    i := len(res) - 1
    for i >= 0 && balance > 0{
        if res[i] == '('{
            balance--
            res = res[:i] + res[i + 1:]
        }
        i--
    }

    return res
}
    