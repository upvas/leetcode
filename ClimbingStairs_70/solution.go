package ClimbingStairs_70

func climbStairs(n int) int {
	dp := map[int]int{}
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; n++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}