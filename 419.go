func countBattleships(board [][]byte) int {
    numberOfBattles := 0
    rows := len(board)
    cols := len(board[0])
    directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    visited := map[int]bool{}
    var dfs func(i, j int)
    dfs = func(i, j int){
        curAdrs := i * cols + j
        if visited[curAdrs]{
            return
        }
        visited[curAdrs] = true
        for _, dir := range directions{
            r := i + dir[0]
            c := j + dir[1]
            if r >= 0 && r < rows && c >= 0 && c < cols && board[r][c] == 'X' && !visited[r * cols + c]{
                dfs(r, c)
            }
        }

    }

    for i := 0; i < rows; i++{
        for j := 0; j < cols; j++{
            if !visited[i * cols + j] && board[i][j] == 'X'{
                numberOfBattles++
                dfs(i, j)
            }
        }
    }
    return numberOfBattles
}