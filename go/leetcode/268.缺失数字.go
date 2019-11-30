/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 *
 * https://leetcode-cn.com/problems/missing-number/description/
 *
 * algorithms
 * Easy (53.17%)
 * Likes:    185
 * Dislikes: 0
 * Total Accepted:    39.5K
 * Total Submissions: 74.2K
 * Testcase Example:  '[3,0,1]'
 *
 * 给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。
 *
 * 示例 1:
 *
 * 输入: [3,0,1]
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: [9,6,4,2,3,5,7,0,1]
 * 输出: 8
 *
 *
 * 说明:
 * 你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?
 *
 */

// package leetcode

// 假设0-n都在序列里面，那么根据求和公式可以算出和
// 另外可以遍历序列求和，两个和之差就是缺失的数字

// @lc code=start
func missingNumber(nums []int) int {
	sum := 0
	// 算实际的和
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	// 根据等差数据求和公式算出1~n的和
	sumN := 0.5 * float32(1+len(nums)) * float32(len(nums))
	return int(sumN) - sum
}

// @lc code=end
