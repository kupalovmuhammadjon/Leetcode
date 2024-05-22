func solve(board [][]byte)  {
    visited := map[int]bool{}
    rows := len(board)
    cols := len(board[0])

    bfs := func(x, y int){
        directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
        q := [][]int{{x, y}}
        visited[x*cols+y] = true
        for len(q) > 0{

            length := len(q)
            for d := 0; d < length; d++{
                i := q[0][0]
                j := q[0][1]
                q = q[1:]
                for _, dir := range directions{
                    r := i + dir[0]
                    c := j + dir[1]
                    if r >= 0 && r < rows && c >= 0 && c < cols && !visited[r*cols+c] && board[r][c] == 'O'{
                        q = append(q, []int{r, c})
                        visited[r*cols+c] = true
                    }
                }
            }
        }

    }

    for i := 0; i < cols; i++ {
        if board[0][i] == 'O' && !visited[0*cols+i] {
            bfs(0, i)
        }
        if board[rows-1][i] == 'O' && !visited[(rows-1)*cols+i] {
            bfs(rows-1, i)
        }
    }

    for i := 0; i < rows; i++ {
        if board[i][0] == 'O' && !visited[i*cols+0] {
            bfs(i, 0)
        }
        if board[i][cols-1] == 'O' && !visited[i*cols+(cols-1)] {
            bfs(i, cols-1)
        }
    }

    for i := 0; i < rows; i++{
        for j := 0; j < cols; j++{
            if !visited[i*cols+j] && board[i][j] == 'O'{
                board[i][j] = 'X'
            }
        }        
    }
}