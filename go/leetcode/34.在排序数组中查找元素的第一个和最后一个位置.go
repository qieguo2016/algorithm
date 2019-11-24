/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (36.82%)
 * Likes:    147
 * Dislikes: 0
 * Total Accepted:    22.4K
 * Total Submissions: 60.9K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 * 
 * 你的算法时间复杂度必须是 O(log n) 级别。
 * 
 * 如果数组中不存在目标值，返回 [-1, -1]。
 * 
 * 示例 1:
 * 
 * 输入: nums = [5,7,7,8,8,10], target = 8
 * 输出: [3,4]
 * 
 * 示例 2:
 * 
 * 输入: nums = [5,7,7,8,8,10], target = 6
 * 输出: [-1,-1]
 * 
 */

// @lc code=start
// 两次二分查找，分别找左右边界
func searchRange(nums []int, target int) []int {
	ret := []int{-1, -1}
	n := len(nums)
	if n <= 0 {
		return ret
	}
	left := 0
	right := n
	// 找左边界，也就是找左侧第一个target的位置，从0往右收敛
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid // 要找左边界，当中间大于等于目标值的时候，左边界在左侧
		}
	}
	if right >= n || nums[right] > target {
		return ret
	}
	ret[0] = right
	// 找右边界，也就是找右侧最后一个target的位置，从end往左收敛
	left = 0
	right = n
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] <= target {  // 在找左边界的时候已经判断了是否有解，所以 right >= target
			left = mid + 1
		} else {
			right = mid
		}
	}
	ret[1] = right - 1
	return ret
}
// @lc code=end

