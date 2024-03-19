func equalPairs(grid [][]int) int {
	count := 0
	cols := [][]int{}
	for i := 0; i < len(grid); i++ {
		col := []int{}
		for j := 0; j < len(grid); j++ {
			col = append(col, grid[j][i])
		}
		cols = append(cols, col)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if isEqual(grid[i], cols[j]) {
				count++
			}
		}
	}

	return count
}
func isEqual(arr1 []int, arr2 []int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}