/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	f1 := 1
	f2 := 2
	for i := 3; i <= n; i++ {
		f2 = f1 + f2  // f2 = f2 + f1
		f1 = f2 - f1  // f1 = f2
	}
	return f2
}

