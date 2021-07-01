package leetcode

func climbStairs(n int) int {
	res:=make(map[int]int)
	res[1] = 1
	res[2] = 2
	for i:=3;i<=n;i++ {
		res[i] = res[i-1]+ res[i-2]
	}
	return res[n]
}
