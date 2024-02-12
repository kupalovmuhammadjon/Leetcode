func finalString(s string) string {

    for i := 0; i < len(s); i++{
        if s[i] == 'i' {
            temp := ""
            for j := i-1; j >= 0; j--{
                temp += string(s[j])
            }
            s = temp + s[i+1:]
            i--
        }
    }

    return s
}