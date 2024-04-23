func minimumTotal(triangle [][]int) int {
    dp := make([][]int, len(triangle))
	// Dynamic Approach 
	// inglizcha yozib ketaman ustoz komentlani
	// first initilize empty dp array
    for i := 0; i < len(triangle); i++{
        dp[i] = make([]int, i+1)
    }
	// only we gonna get last row od values which means it will be a bottom-up approach
    lastRowIndex := len(triangle) - 1
    for i, val := range triangle[lastRowIndex] {
        dp[lastRowIndex][i] = val
    }
	// every cell of dp contains minimum sum required to come to that point
    for i := len(triangle) - 2; i >= 0; i-- {
        for j := 0; j < len(triangle[i]); j++ {
            dp[i][j] = triangle[i][j] + min(dp[i+1][j], dp[i+1][j+1])
        }
    }
    return dp[0][0]
}
func min(x, y int) int{
    if x > y{
        return y
    }
    return x
}