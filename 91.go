func numDecodings(s string) int {
    chars := map[string]string{}
    for i := 0; i < 26; i++{
        chars[strconv.Itoa(i+1)] = string(byte(i + 'A'))
    }
    memo := map[int]int{}
    var backtrack func(int) int
    backtrack = func(i int) int{
        if i >= len(s){
            return 1
        }
        if _, ok := memo[i]; ok{
            return memo[i]
        }
        count := 0
       
        if _, ok := chars[s[i:i+1]]; ok{
            count += backtrack(i+1)
        }
        if i + 2 <= len(s){
            if _, ok := chars[s[i:i+2]]; ok{
                count += backtrack(i+2)
            }
        }
        memo[i] = count
        return count
    }
    
    return backtrack(0)
}