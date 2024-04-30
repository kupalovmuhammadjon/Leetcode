func letterCombinations(digits string) []string {
    combinations := []string{}
    if len(digits) == 0{
        return combinations
    }
    keyboard := map[byte][]string{'2':[]string{"a", "b", "c"},
    '3':[]string{"d", "e", "f"}, '4':[]string{"g", "h", "i"},
    '5':[]string{"j", "k", "l"}, '6':[]string{"m", "n", "o"},
    '7':[]string{"p", "q", "r", "s"}, '8':[]string{"t", "u", "v"},
    '9':[]string{"w", "x", "y", "z"}}

    var backtrack func(i int, currentComb string)
    backtrack = func(i int, currentComb string){
        if len(currentComb) == len(digits){
            combinations = append(combinations, currentComb)
            return
        }
        for _, v := range keyboard[digits[i]]{
            backtrack(i+1, currentComb+v)
        }
    }
    backtrack(0, "")

    return combinations
}