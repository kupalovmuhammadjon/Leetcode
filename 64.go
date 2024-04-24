func minPathSum(grid [][]int) int {
    // save minimum path of every cell and at very end answer come
    for i := 0; i < len(grid); i++{
        for j := 0; j < len(grid[i]); j++{
            left := math.MaxInt64
            right := math.MaxInt64
            if j - 1 >= 0 {
                left = grid[i][j-1]
            }
            if i - 1 >= 0{
                right = grid[i-1][j]
            }
            if right != math.MaxInt64 || left != math.MaxInt64{
                grid[i][j] += min(right, left)
            }
        }
    } 
    return grid[len(grid)-1][len(grid[0])-1]
}