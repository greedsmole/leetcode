func calcDp(dp []int, n int) int {
	if n <= 1 {
		dp[n] = 1
		return dp[n]
	}
	if dp[n] != -1 {
		return dp[n]
	}
	dp[n] = calcDp(dp, n-1) + calcDp(dp, n-2)
	return dp[n]
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}
	return calcDp(dp, n)
}