/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 *
 * https://leetcode-cn.com/problems/find-the-duplicate-number/description/
 *
 * algorithms
 * Medium (61.75%)
 * Likes:    328
 * Dislikes: 0
 * Total Accepted:    26.4K
 * Total Submissions: 42.8K
 * Testcase Example:  '[1,3,4,2,2]'
 *
 * 给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和
 * n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
 * 
 * 示例 1:
 * 
 * 输入: [1,3,4,2,2]
 * 输出: 2
 * 
 * 
 * 示例 2:
 * 
 * 输入: [3,1,3,4,2]
 * 输出: 3
 * 
 * 
 * 说明：
 * 
 * 
 * 不能更改原数组（假设数组是只读的）。
 * 只能使用额外的 O(1) 的空间。
 * 时间复杂度小于 O(n^2) 。
 * 数组中只有一个重复的数字，但它可能不止重复出现一次。
 * 
 * 
 */
// 题目中的数字都是1~n，假设没有重复，那按中点分割每一侧的数字个数=max-min
// 如果存在重复，那么有n>max-min，不断二分这个mid可以找到目标值

// @lc code=start
func findDuplicate(nums []int) int {
	left, right := 1, len(nums)
	for left < right {
		mid := left + (right - left) / 2
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count <= mid {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
// @lc code=end

