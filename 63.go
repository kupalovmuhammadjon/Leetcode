func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if obstacleGrid[0][0] == 1{
        return 0
    }
    memo := map[int]int{}
    rows := len(obstacleGrid)
    cols := len(obstacleGrid[0])
    var backtrack func(i, j int) int
    backtrack = func(i, j int) int{
        curAdrs := i * cols + j
        if i == rows-1 && j == cols-1{
            return 1
        }
        if _, ok := memo[curAdrs]; ok{
            return memo[curAdrs]
        }
        count := 0
        if i + 1 < rows && obstacleGrid[i+1][j] != 1{
            count += backtrack(i+1, j)
        }
        if j + 1 < cols && obstacleGrid[i][j+1] != 1{
            count += backtrack(i, j+1)
        }
        memo[curAdrs] = count
        return count
    }

    return backtrack(0, 0)
}