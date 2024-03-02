func transpose(matrix [][]int) [][]int {
	res := [][]int{}
	for i := 0; i < len(matrix[0]); i++ {
		temp := []int{}
		for j := 0; j < len(matrix); j++ {
			temp = append(temp, matrix[j][i])
		}
		res = append(res, temp)
	}
	return res
}