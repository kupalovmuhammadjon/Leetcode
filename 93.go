func restoreIpAddresses(s string) []string {
    validIps := []string{}

    var backtrack func(curIp string, i, count int)
    backtrack = func(curIp string, i, count int){
        if count == 4 && i == len(s){
            validIps = append(validIps, curIp)
            return
        }
        if count > 4{
            return
        }
        for l := 1; l <= 3 && l + i <= len(s); l++{
            part := s[i:l+i]
            if isValidPart(part){
                if count == 0{
                    backtrack(curIp+part, i+l, count+1)
                }else{
                    backtrack(curIp+"."+part, i+l, count+1)
                }
            }
        }

    }
    backtrack("", 0, 0)

    return validIps
}
func isValidPart(part string) bool{
    if part[0] == '0' && len(part) > 1{
        return false
    }

    n, err := strconv.Atoi(part)
    if err != nil || n > 255{
        return false 
    }
    return true
}