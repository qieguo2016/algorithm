/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 *
 * https://leetcode-cn.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (35.91%)
 * Likes:    308
 * Dislikes: 0
 * Total Accepted:    26.1K
 * Total Submissions: 72.4K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。
 *
 * 示例 1:
 *
 * 输入: [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 *
 *
 * 示例 2:
 *
 * 输入: [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 *
 */

// package leetcode

// 跟53.最大子序和有点类似，但是乘法中会有负数存在，而且负负可以得正
// 所以可以用max和min分别记录正向和负向的最大值，每次判断一下当前数的正负区别对待
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

func maxProduct(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	res := nums[0]
	maximum, minimum := res, res
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		if cur > 0 {
			minimum = min(minimum*cur, cur)
			maximum = max(maximum*cur, cur)
		} else {
			tmp := minimum
			minimum = min(maximum*cur, cur)
			maximum = max(tmp*cur, cur)
		}
		res = max(res, maximum)
	}
	return res
}

// @lc code=end
