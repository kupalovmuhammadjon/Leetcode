func appendCharacters(s string, t string) int {
    tInd := 0
    for i := 0; i < len(s) && tInd < len(t); i++{
        if s[i] == t[tInd]{
            tInd++
        }
    }

    return len(t) - tInd
}