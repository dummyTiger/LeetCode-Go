package leetcode

func fib(n int) int {
	dp := make(map[int]int)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}
