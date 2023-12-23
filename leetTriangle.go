package leetTriangle

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([]int, m+1)
	var i, j int
	for i = m - 1; i >= 0; i-- {
		for j = 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}
