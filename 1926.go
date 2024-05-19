func nearestExit(maze [][]byte, entrance []int) int {
    
    level := 0
    maze[entrance[0]][entrance[1]] = '+'
    q := [][]int{entrance}
    directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    for len(q) > 0{
        level++
        length := len(q)
        for k := 0; k < length; k++{
            i := q[0][0]
            j := q[0][1]
            q = q[1:]
            
            for _, dir := range directions{
                newI := i + dir[0]
                newJ := j + dir[1]
                if newI >= 0 && newI < len(maze) && newJ >= 0 && newJ < len(maze[0]) && maze[newI][newJ] == '.'{
                    if newI == len(maze)-1 || newJ == len(maze[0])-1 || newI == 0 || newJ == 0 {
                        return level
                    }
                    q = append(q, []int{newI, newJ})
                    maze[newI][newJ] = '+'
                }
            }

        }
    }

    return -1
}