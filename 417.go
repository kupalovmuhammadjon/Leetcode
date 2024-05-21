func pacificAtlantic(heights [][]int) [][]int {
    coordinates := [][]int{}
    rows := len(heights)
    cols := len(heights[0])
    memo := map[int][2]bool{}
    directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    var dfs func(i, j, node int, visited map[int]bool) [2]bool
    dfs = func(i, j, node int, visited map[int]bool) [2]bool{
        canReach := [2]bool{} 
        curAdrs := i * cols + j
        // if r, ok := memo[curAdrs]; ok{
        //     return r
        // }
        if memo[curAdrs][0]{
            canReach[0] = true
        }
        if memo[curAdrs][1]{
            canReach[1] = true
        }
        if i == 0 || j == 0{
            canReach[0] = true
        }
        if i == rows-1 || j == cols-1{
            canReach[1] = true
        }
        if canReach[0] && canReach[1]{
            return canReach
        }
        visited[i*cols+j] = true
        for _, direction := range directions{
            r := i + direction[0]
            c := j + direction[1]
            if r >= 0 && r < rows && c >= 0 && c < cols && !visited[r*cols+c] && node >= heights[r][c]{
                reachable := dfs(r, c, heights[r][c], visited)
                // o'zgrtirish natija bilan
                canReach[0] = canReach[0] || reachable[0]
                canReach[1] = canReach[1] || reachable[1]
            }
        }

        memo[curAdrs] = canReach
        visited[curAdrs] = false

        return canReach
    }

    for i := 0; i < rows; i++{
        for j := 0; j < cols; j++{
            r := dfs(i, j, heights[i][j], map[int]bool{})
            if r[0] && r[1]{
                coordinates = append(coordinates, []int{i, j})
            }
        }
    }

    return coordinates
}