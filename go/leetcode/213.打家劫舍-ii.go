/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 *
 * https://leetcode-cn.com/problems/house-robber-ii/description/
 *
 * algorithms
 * Medium (35.22%)
 * Likes:    146
 * Dislikes: 0
 * Total Accepted:    14.8K
 * Total Submissions: 41.8K
 * Testcase Example:  '[2,3,2]'
 *
 *
 * 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
 *
 * 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
 *
 * 示例 1:
 *
 * 输入: [2,3,2]
 * 输出: 3
 * 解释: 你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
 *
 *
 * 示例 2:
 *
 * 输入: [1,2,3,1]
 * 输出: 4
 * 解释: 你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
 * 偷窃到的最高金额 = 1 + 3 = 4 。
 *
 */

// package leetcode
// 思路跟198一样，关键是将环换成普通队列形式，那么可以将头尾去掉，然后再比较头尾
// 可以分成[0, n-1]和[1, n]两种组合，分别覆盖了头和尾部，取两者大值即可

// @lc code=start

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	if len(nums) <= 1 {
		return nums[0]
	}
	return max(helper(nums, 0, len(nums)-1), helper(nums, 1, len(nums)))
}

func helper(nums []int, left int, right int) int {
	even, odd := 0, 0
	for i := left; i < right; i++ {
		cur := nums[i]
		if i%2 == 0 {
			even = max(even+cur, odd)
		} else {
			odd = max(odd+cur, even)
		}
	}
	return max(even, odd)
}

// @lc code=end
