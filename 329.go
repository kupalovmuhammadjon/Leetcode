func longestIncreasingPath(matrix [][]int) int {
    directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    memo := map[int]int{}
    var dfs func(i, j, prev int, visited map[int]bool) int  
    dfs = func(i, j, prev int, visited map[int]bool) int {
        if memo[i*len(matrix[0])+j] > 0{
            return memo[i*len(matrix[0])+j]
        }
        visited[i*len(matrix[0])+j] = true
        longest := 1 
        for _, dir := range directions {
            r := i + dir[0]
            c := j + dir[1]
            if r >= 0 && r < len(matrix) && c >= 0 && c < len(matrix[0]) && !visited[r*len(matrix[0])+c] && matrix[r][c] > prev {
                longest = max(longest, 1+dfs(r, c, matrix[r][c], visited)) 
            }
        }
        visited[i*len(matrix[0])+j] = false
        memo[i*len(matrix[0])+j] = longest
        return longest
    }

    longestInc := -1
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            mx := dfs(i, j, matrix[i][j], map[int]bool{})
            if longestInc < mx {
                longestInc = mx
            }
        }
    }
    return longestInc
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
