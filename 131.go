func partition(s string) [][]string {
    partitions := [][]string{}
    
    var backtrack func(currentCombo []string, ind int)
    backtrack = func(currentCombo []string, ind int){
        if ind == len(s){
            cpy := make([]string, len(currentCombo))
            copy(cpy, currentCombo)
            partitions = append(partitions, cpy)
            return
        }

        for i := ind; i < len(s); i++{
            if isPalindrome(s[ind:i+1]){
                currentCombo = append(currentCombo, s[ind:i+1])
                backtrack(currentCombo, i+1)
                currentCombo = currentCombo[:len(currentCombo)-1]
            }
        }
    }
    backtrack([]string{}, 0)
    return partitions
}
func isPalindrome(s string) bool{
    for i := 0; i < len(s)/2; i++{
        if s[i] != s[len(s)-1-i]{
            return false
        }
    }
    return true
}