func letterCasePermutation(s string) []string {
    permutations := []string{}

    var backtrack func(curStr string, i int) 
    backtrack = func(curStr string, i int) {
        if len(curStr) == len(s){
            permutations = append(permutations, curStr)
            return
        }

        curStr += string(s[i])
        backtrack(curStr, i+1)
        curStr = curStr[:len(curStr)-1]
        if s[i] >= 'a' && s[i] <= 'z' {
            curStr += string(s[i]-32)
            backtrack(curStr, i+1)
        }else if s[i] >= 'A' && s[i] <= 'Z' {
            curStr += string(s[i]+32)
            backtrack(curStr, i+1)
        }
    }
    backtrack("", 0)
    return permutations
}