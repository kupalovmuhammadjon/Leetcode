func islandPerimeter(grid [][]int) int {
    perimetr := 0

    for i := 0; i < len(grid); i++{
        for j := 0; j < len(grid[0]); j++{
            if grid[i][j] == 1{
                perimetr += calculateSurroundingWaterArea(i, j, grid)
            }
        } 
    }
    return perimetr
}
func calculateSurroundingWaterArea(i, j int, grid [][]int) int{
    totalBeachAreaofCube := 0

    if i - 1 < 0 || grid[i-1][j] == 0{
        totalBeachAreaofCube++
    }
    if i + 1 >= len(grid) || grid[i+1][j] == 0{
        totalBeachAreaofCube++
    }
    if j - 1 < 0 || grid[i][j-1] == 0{
        totalBeachAreaofCube++
    }
    if j + 1 >= len(grid[0]) || grid[i][j+1] == 0{
        totalBeachAreaofCube++
    }
    return totalBeachAreaofCube
}