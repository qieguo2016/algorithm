/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (52.07%)
 * Likes:    1695
 * Dislikes: 0
 * Total Accepted:    462.3K
 * Total Submissions: 888K
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 *
 * 注意：给定 n 是一个正整数。
 *
 * 示例 1：
 *
 * 输入： 2
 * 输出： 2
 * 解释： 有两种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶
 * 2.  2 阶
 *
 * 示例 2：
 *
 * 输入： 3
 * 输出： 3
 * 解释： 有三种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶 + 1 阶
 * 2.  1 阶 + 2 阶
 * 3.  2 阶 + 1 阶
 *
 *
 */

// @lc code=start
func climbStairs(n int) int {
	// 动态规划，最后一步要么走1阶，要么走2阶
	// f(n) = f(n-1) + f(n-2)
	// 推广到一次最多可走k阶，那 f(n) = sum(f(n-i)), 其中i属于[1,k]
	if n <= 2 {
		return n
	}
	// dp := make([]int, n)
	// dp[0] = 1
	// dp[1] = 2
	// for i := 2; i < n; i++ {
	// 	dp[i] = dp[i-1] + dp[i-2]
	// }
	// return dp[n-1]

	// 上面的空间复杂度是O(n)，但每次只有两个变量，所以用两个变量来代替
	dp2 := 1 // n-2
	dp1 := 2 // n-1
	dpn := 0
	for i := 2; i < n; i++ {
		dpn = dp1 + dp2
		dp2 = dp1
		dp1 = dpn
	}
	return dpn
}

// @lc code=end

