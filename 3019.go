func countKeyChanges(s string) int {
    st := strings.ToLower(s)
    count := 0
    for i := 1; i < len(st); i++{
        if st[i-1] != st[i]{
            count++
        }
    } 
    return count
}