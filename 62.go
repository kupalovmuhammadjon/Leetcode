func uniquePaths(m int, n int) int {
    
    memo := map[int]int{}
    var backtrack func(i, j int) int
    backtrack = func(i, j int) int{
        curAdrs := i * n + j
        if i == m-1 && j == n-1{
            return 1
        }
        if _, ok := memo[curAdrs]; ok{
            return memo[curAdrs]
        }
        count := 0
        if i + 1 < m{
            count += backtrack(i+1, j)
        }
        if j + 1 < n{
            count += backtrack(i, j+1)
        }
        memo[curAdrs] = count
        return count
    }

    return backtrack(0, 0)
}