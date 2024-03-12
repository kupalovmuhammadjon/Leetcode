func updateMatrix(mat [][]int) [][]int {

	zeros := [][]int{}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				cor := []int{i, j}
				zeros = append(zeros, cor)
			}
		}
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				mat[i][j] = nearZero(zeros, i, j)
			}
		}
	}

	return mat
}
func nearZero(zeros [][]int, i int, j int) int {

	mn := 12323
	for k := 0; k < len(zeros); k++ {
		dis := abs(zeros[k][0]-i) + abs(zeros[k][1]-j)
		if mn > dis {
			mn = dis
		}
	}

	return mn
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}