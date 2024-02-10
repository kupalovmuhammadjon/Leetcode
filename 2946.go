func areSimilar(mat [][]int, k int) bool {
    rows := len(mat)
    cols := len(mat[0])
    cpy := make([][]int, rows)
    for i := 0; i < rows; i++{
        cpy[i] = make([]int, cols)
        copy(cpy[i], mat[i])
    }
    for c := 0; c < k; c++ {
        for i := 0; i < len(mat); i++{
            d := mat[i][0]
            row := mat[i][1:]
            row = append(row, d)
            mat[i] = row
        }
    }
    for i := 0; i < len(mat); i++{
        for j := 0; j < len(mat[0]); j++{
            if cpy[i][j] != mat[i][j]{
                return false
            }
        }
    }
    return true
}