/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 *
 * https://leetcode-cn.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (55.98%)
 * Likes:    330
 * Dislikes: 0
 * Total Accepted:    41.8K
 * Total Submissions: 73.9K
 * Testcase Example:  '3\n2'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 * 
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 * 
 * 问总共有多少条不同的路径？
 * 
 * 
 * 
 * 例如，上图是一个7 x 3 的网格。有多少可能的路径？
 * 
 * 说明：m 和 n 的值均不超过 100。
 * 
 * 示例 1:
 * 
 * 输入: m = 3, n = 2
 * 输出: 3
 * 解释:
 * 从左上角开始，总共有 3 条路径可以到达右下角。
 * 1. 向右 -> 向右 -> 向下
 * 2. 向右 -> 向下 -> 向右
 * 3. 向下 -> 向右 -> 向右
 * 
 * 
 * 示例 2:
 * 
 * 输入: m = 7, n = 3
 * 输出: 28
 * 
 */
func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	// paths[i][j] = paths[i-1][j] + paths[i][j-1]
	// 可换成1维度数组，只留下列，可以发现dp[j] = dp[j] + dp[j-1]
	// 1, 1, 1, 1, 1, 1, 1
	// 1, 2, 3, 4, 5, 6, 7
	// 1, 3, 6, 10,15,21,28
	dp := make([]int, n)
	for i := 0; i < n; i++ { // 第1行全是1
		dp[i] = 1
	}
	for i := 1; i < m; i++ { 	// 第1行全是1
		for j := 1; j < n; j++ {  // 第一列全是1
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}
// 题目是一个组合问题，机器人必然要向下走m-1步，向右走n-1步
// 换个角度看，机器人必定要走m+n-2步，其中选出m-1步往下走，这就变成了一个组合数问题
// 假设m<=n，那么结果就是C(m-1,m+n-2)，而C(m, n)=n!/(m!*(n-m)!)

