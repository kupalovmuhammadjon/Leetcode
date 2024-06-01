func scoreOfString(s string) int {
    score := 0
    for i := 1; i < len(s); i++{
        score += abs(int(s[i-1]) - int(s[i]))
    }
    return score
}

func abs(x int) int{
    if x < 0{
        return -x
    }
    return x
}