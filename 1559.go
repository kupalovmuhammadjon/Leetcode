func containsCycle(grid [][]byte) bool {
    rows := len(grid)
    cols := len(grid[0])
    directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    visited := map[int]bool{}

    var bfs func(x, y int) bool
    bfs = func(x, y int) bool {
        q := [][]int{{x, y, -1, -1}}
        visited[x*cols+y] = true
        for len(q) > 0 {
            i := q[0][0]
            j := q[0][1]
            pi := q[0][2]
            pj := q[0][3]
            q = q[1:]
            for _, dir := range directions {
                r := i + dir[0]
                c := j + dir[1]
                if r >= 0 && r < rows && c >= 0 && c < cols && grid[i][j] == grid[r][c] {
                    if visited[r*cols+c] {
                        if r != pi || c != pj {
                            return true
                        }
                    } else {
                        q = append(q, []int{r, c, i, j})
                        visited[r*cols+c] = true
                    }
                }
            }
        }
        return false
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if !visited[i*cols+j] {
                if bfs(i, j) {
                    return true
                }
            }
        }
    }
    return false
}
