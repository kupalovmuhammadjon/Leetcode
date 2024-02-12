func findChampion(grid [][]int) int {
	pos := 0
	mxc := 0
	for i := 0; i < len(grid); i++ {
		count := 0
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
		if mxc < count {
			mxc = count
			pos = i
		}
	}
	return pos
}