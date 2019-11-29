/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 *
 * https://leetcode-cn.com/problems/house-robber/description/
 *
 * algorithms
 * Easy (42.16%)
 * Likes:    546
 * Dislikes: 0
 * Total Accepted:    57.7K
 * Total Submissions: 136.8K
 * Testcase Example:  '[1,2,3,1]'
 *
 *
 * 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 *
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
 *
 * 示例 1:
 *
 * 输入: [1,2,3,1]
 * 输出: 4
 * 解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 *
 * 示例 2:
 *
 * 输入: [2,7,9,3,1]
 * 输出: 12
 * 解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
 * 偷窃到的最高金额 = 2 + 9 + 1 = 12 。
 *
 *
 */
// package leetcode

// @lc code=start

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	even, odd := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			even = max(even+nums[i], odd) // 偷偶数的话要看跟奇数的比谁大
		} else {
			odd = max(odd+nums[i], even)
		}
	}
	return max(even, odd)
}

// 另外一种方法是用普通dp，dp[i]=max(dp[i-2]+nums[i], dp[i-1])，
// 有两种选择，偷则加上上上一个，不偷则是取上一个，最后值是取两种选择中的更大值
// func rob(nums []int) int {
// 	l := len(nums)
// 	if l <= 0 {
// 		return 0
// 	}
// 	if l <= 1 {
// 		return nums[0]
// 	}
// 	dp := make([]int, l)
// 	dp[0] = nums[0]
// 	dp[1] = max(nums[0], nums[1])
// 	for i := 2; i < l; i++ {
// 		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
// 	}
// 	return dp[l-1]
// }

// @lc code=end
