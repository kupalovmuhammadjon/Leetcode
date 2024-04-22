func maximumHappinessSum(happiness []int, k int) int64 {
    sort.Ints(happiness)
    totalHappiness := 0    
    j := 0
    for i := len(happiness) - 1; i  >= len(happiness) - k ; i--{
        if happiness[i] - j > 0{
            totalHappiness += happiness[i] - j
        }
        j++
    }

    return int64(totalHappiness)
}