func shortestPathBinaryMatrix(grid [][]int) int {
    if grid[0][0] == 1 || grid[len(grid)-1][len(grid)-1] == 1{
        return -1
    }
    minimumPath := 0
    neighbours := [][]int{{1, 0}, {0, 1}, {1, 1}, {-1, 0}, {0, -1}, {-1, -1}, {-1, 1}, {1, -1}}
    q := [][]int{{0, 0}}
    for len(q) > 0{
        minimumPath++
        length := len(q)
        for k := 0; k < length; k++{
            i := q[0][0]
            j := q[0][1]
            if i == len(grid)-1 && j == len(grid)-1{
                return minimumPath
            }
            q = q[1:]
            for _, dir := range neighbours{
                newI := i + dir[0]
                newJ := j + dir[1]
                if inBounds(len(grid), newI, newJ) && grid[newI][newJ] == 0{
                    q = append(q, []int{newI, newJ})
                    grid[newI][newJ] = 1
                }
            }
        }
    }
    return -1
}

func inBounds(length, i, j int) bool{
    return i >= 0 && j >= 0 && i < length && j < length
}