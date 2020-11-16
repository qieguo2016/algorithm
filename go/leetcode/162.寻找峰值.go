/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 *
 * https://leetcode-cn.com/problems/find-peak-element/description/
 *
 * algorithms
 * Medium (47.80%)
 * Likes:    327
 * Dislikes: 0
 * Total Accepted:    65.5K
 * Total Submissions: 136.7K
 * Testcase Example:  '[1,2,3,1]'
 *
 * 峰值元素是指其值大于左右相邻值的元素。
 * 
 * 给定一个输入数组 nums，其中 nums[i] ≠ nums[i+1]，找到峰值元素并返回其索引。
 * 
 * 数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。
 * 
 * 你可以假设 nums[-1] = nums[n] = -∞。
 * 
 * 示例 1:
 * 
 * 输入: nums = [1,2,3,1]
 * 输出: 2
 * 解释: 3 是峰值元素，你的函数应该返回其索引 2。
 * 
 * 示例 2:
 * 
 * 输入: nums = [1,2,1,3,5,6,4]
 * 输出: 1 或 5 
 * 解释: 你的函数可以返回索引 1，其峰值元素为 2；
 * 或者返回索引 5， 其峰值元素为 6。
 * 
 * 
 * 说明:
 * 
 * 你的解法应该是 O(logN) 时间复杂度的。
 * 
 */

// @lc code=start
// 好好利用左右边界都是负无穷这个条件，二分查找
// 向上趋势则右侧必有峰值，向下趋势则左侧必有峰值
func findPeakElement(nums []int) int {
	return findHelper(nums, 0, len(nums)-1)
}

func findHelper(nums []int, l, r int) int {
	if l == r {
		return l
	}
	mid := l + (r - l) / 2
	if nums[mid]>nums[mid+1] {
		return findHelper(nums, l, mid)
	}
	return findHelper(nums, mid+1, r)
}
// @lc code=end

