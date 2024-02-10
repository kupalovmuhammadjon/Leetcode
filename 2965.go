func findMissingAndRepeatedValues(grid [][]int) []int {

    ls := []int{}

    for i := 0; i < len(grid); i++{
        for j := 0; j < len(grid[0]); j++{
            ls = append(ls, grid[i][j])
        }
    }
    slices.Sort(ls)
    res := []int{}
    for i := 0; i < len(ls); i++{
        for j := i + 1; j < len(ls); j++{
            if ls[i] == ls[j]{
                res = append(res, ls[j])
                break
            }
        }
    }
    for i := 1; i < len(ls)+1; i++{
        if ! isInit(ls, i){
            res = append(res, i)
        }
    }

    return res
}
func isInit(arr []int, t int) bool{
    for i:= 0; i < len(arr); i++{
        if t == arr[i]{
            return true
        }
    }
    return false
}