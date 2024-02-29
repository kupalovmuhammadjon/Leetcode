func numMagicSquaresInside(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[0])-2; j++ {
			if isMagic(grid, i, j) {
				count++
			}
		}
	}
	return count
}

func isMagic(grid [][]int, i int, j int) bool {
	nums := make(map[int]bool)
	sum := grid[i][j] + grid[i][j+1] + grid[i][j+2]

	for k := i; k < i+3; k++ {
		smr := 0
		for h := j; h < j+3; h++ {
			nums[grid[k][h]] = true
			smr += grid[k][h]
		}
		if smr != sum {
			return false
		}
	}

	for k := j; k < j+3; k++ {
		smc := 0
		for h := i; h < i+3; h++ {
			smc += grid[h][k]
		}
		if smc != sum {
			return false
		}
	}

	d1 := grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2]
	d2 := grid[i][j+2] + grid[i+1][j+1] + grid[i+2][j]

	if d1 != sum || d2 != sum {
		return false
	}

	for num := 1; num <= 9; num++ {
		if !nums[num] {
			return false
		}
	}

	return true
}
