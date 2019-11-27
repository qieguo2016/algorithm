/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 *
 * https://leetcode-cn.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (36.38%)
 * Likes:    280
 * Dislikes: 0
 * Total Accepted:    30.2K
 * Total Submissions: 83K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给定不同面额的硬币 coins 和一个总金额
 * amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
 * 
 * 示例 1:
 * 
 * 输入: coins = [1, 2, 5], amount = 11
 * 输出: 3 
 * 解释: 11 = 5 + 5 + 1
 * 
 * 示例 2:
 * 
 * 输入: coins = [2], amount = 3
 * 输出: -1
 * 
 * 说明:
 * 你可以认为每种硬币的数量是无限的。
 * 
 */

// 求极值一般会考虑使用动态规划。用dp[i]表示i元的最少硬币兑换数，则状态转移方程：
// dp[i] = min(dp[i], dp[i-coins[x]]+1)
// dp有两种模式，一种在迭代的时候增加记忆数组，另外一种是上来先算记忆数组
// PS: 这个兑换有个特例，就是各种硬币是倍数关系的时候，这个时候退化成贪婪模式

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		// 假设硬币最小是1元，所以最大是amount个硬币
		// 但是现在不一定有1元硬币，所以初值设置成amount+1就可以保证进入到-1的分支
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
// @lc code=end

