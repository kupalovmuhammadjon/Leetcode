func orangesRotting(grid [][]int) int {

    bfs := func(queue [][]int) (int, int)  {
        
        rotted := 0
        minTime := -1
        neighbours := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

        for len(queue) > 0{
            minTime++
            length := len(queue)

            for h := 0; h < length; h++{
                curI := queue[0][0]
                curJ := queue[0][1]
                queue = queue[1:]

                for _, dir := range neighbours{
                    newI := curI + dir[0]
                    newJ := curJ + dir[1]

                    if isNotOutOfBounds(len(grid), len(grid[0]), newI, newJ) && grid[newI][newJ] == 1{
                        grid[newI][newJ] = 2
                        rotted++
                        queue = append(queue, []int{newI, newJ})
                    }
                }
            }
        }
        return rotted, minTime
    }

    var freshOranges int
    queue := [][]int{}
    
    for i := 0; i < len(grid); i++{
        for j := 0; j < len(grid[i]); j++{
            if grid[i][j] == 2{
                queue = append(queue, []int{i, j})
            }else if grid[i][j] == 1{
                freshOranges++
            }
        }
    }

    if freshOranges == 0{
        return 0
    }
    rotted, time := bfs(queue)
    fmt.Println(freshOranges, rotted)
    if freshOranges != rotted{
        return -1
    }

    return time
}

func isNotOutOfBounds(row, col, i, j int) bool {
    return i >= 0 && j >= 0 && i < row && j < col
}
