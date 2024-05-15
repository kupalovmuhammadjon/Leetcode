func wordBreak(s string, wordDict []string) []string {
    sentences := []string{}
    words := map[string]bool{}
    for _, word := range wordDict{
        words[word] = true
    }

    var backtrack func(start int, currentSent string)
    backtrack = func(start int, currentSent string){
        if start == len(s){
            currentSent = currentSent[:len(currentSent)-1]
            sentences = append(sentences, currentSent)
            return
        }

        for i := start; i < len(s); i++{
            if words[s[start:i+1]]{
                // currentSent += s[start:i+1] + " "
                backtrack(i+1, currentSent + s[start:i+1] + " ")
            }
        }
    }
    backtrack(0, "")

    return sentences
}