/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 *
 * https://leetcode-cn.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (43.34%)
 * Likes:    360
 * Dislikes: 0
 * Total Accepted:    38.6K
 * Total Submissions: 88.7K
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * 给定一个无序的整数数组，找到其中最长上升子序列的长度。
 *
 * 示例:
 *
 * 输入: [10,9,2,5,3,7,101,18]
 * 输出: 4
 * 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
 *
 * 说明:
 *
 *
 * 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
 * 你算法的时间复杂度应该为 O(n^2) 。
 *
 *
 * 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
 *
 */

/*
 * 求极值第一反应是动态规划，状态转移方程是dp[i]=max(dp[i], nums[i]>nums[j] ? dp[j]+1 : dp[j])
 * 其中，dp[i]表示0~i之间最长上升子序列长，0<j<i；遍历0-i，如果小于nums[i]则与dp[i]对比一下，看看能否更新
 * PS: 如果要求连续的最长则可以简化成只记录一个max变量，然后找各个子串的最长值
 */

// package leetcode

// @lc code=start
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	res := 1
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1 // 每个元素不与别的元素比较最短都是1个长度
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
