package main


// 斐波那契数列
func fib(n int) int {
	if n < 2{
		return n
	}
	//p,q,r := 0,0,1
	//for i := 2; i <= n; i++ {
	//	q = p
	//	p = r
	//	r = q+p
	//}
	//return r

	dp := make([]int,n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <=n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {

}
