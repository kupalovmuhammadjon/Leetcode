func uniquePathsIII(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	si := 0
	sj := 0
	emptyCells := 2

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				emptyCells++
			} else if grid[i][j] == 1 {
				si = i
				sj = j
			}
		}
	}
	numberOfPaths := 0
	visited := map[int]bool{}
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var backtrack func(i, j, numOfVisitedCells int)
	backtrack = func(i, j, numOfVisitedCells int) {
		visited[i*cols+j] = true
		if grid[i][j] == 2 && numOfVisitedCells == emptyCells {
            visited[i*cols+j] = false
			numberOfPaths++
			return
		}
		for _, dir := range directions {
			r := i + dir[0]
			c := j + dir[1]
			if r >= 0 && c >= 0 && r < rows && c < cols && !visited[r*cols+c] && grid[r][c] != -1 {
				backtrack(r, c, numOfVisitedCells+1)
			}
		}
		visited[i*cols+j] = false
	}

	backtrack(si, sj, 1)

	return numberOfPaths
}