func getMaximumGold(grid [][]int) int {
    maxPath := 0
    visited :=  map[int]bool{}
    var backtrack func(currentSum int, i, j int)
    backtrack = func(currentSum int, i, j int){
        visited[i*len(grid[i])+j] = true
        currentSum += grid[i][j]
        if maxPath < currentSum{
            maxPath = currentSum
        }
        if i - 1 >= 0 && grid[i-1][j] != 0 && !visited[(i-1)*len(grid[i])+j]{
            backtrack(currentSum, i-1, j)
        }

        if i + 1 < len(grid) && grid[i+1][j] != 0 && !visited[(i+1)*len(grid[i])+j]{
            backtrack(currentSum, i+1, j)
        }

        if j - 1 >= 0 && grid[i][j-1] != 0 && !visited[i*len(grid[i])+j-1]{
            backtrack(currentSum, i, j-1)
        }

        if j + 1 < len(grid[i]) && grid[i][j+1] != 0 && !visited[i*len(grid[i])+j+1]{
            backtrack(currentSum, i, j+1)
        }
        delete(visited, i*len(grid[i])+j)
    }
    for i := 0; i < len(grid); i++{
        for j := 0; j < len(grid[i]); j++{
            backtrack(0, i, j)
        }
    }

    return maxPath
}