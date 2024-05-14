func exist(board [][]byte, word string) bool {
    visited := map[int]bool{}
    chars := map[byte]bool{}
    for i := 0; i < len(word); i++{
        chars[word[i]] = true
    }
    directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    res := false
    var backtrack func(curStr int, i, j int)
    backtrack = func(curStr int, i, j int){
        if curStr == len(word){
            res = true
            return 
        }
        if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || word[curStr] != board[i][j] || visited[i * len(board[i])+j]{
            return
        }
        visited[i * len(board[i])+j] = true
        for _, dir := range directions{
            backtrack(curStr+1,i+dir[0],j+dir[1])
        } 
        delete(visited, i * len(board[i])+j)
    }

    for i := 0; i < len(board); i++{
        for j := 0; j < len(board[i]); j++{
            if chars[board[i][j]] && word[0] == board[i][j]{
                backtrack(0, i, j) 
            }
        }
    }
    return res
}