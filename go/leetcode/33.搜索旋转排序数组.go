/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (36.68%)
 * Likes:    253
 * Dislikes: 0
 * Total Accepted:    25.8K
 * Total Submissions: 70.4K
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 * 
 * ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
 * 
 * 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
 * 
 * 你可以假设数组中不存在重复的元素。
 * 
 * 你的算法时间复杂度必须是 O(log n) 级别。
 * 
 * 示例 1:
 * 
 * 输入: nums = [4,5,6,7,0,1,2], target = 0
 * 输出: 4
 * 
 * 
 * 示例 2:
 * 
 * 输入: nums = [4,5,6,7,0,1,2], target = 3
 * 输出: -1
 * 
 */

// O(log n)级别的时间复杂度，第一反应就是二分查找了
// 因为有序数组在某点上旋转了，所以二分查找还需要根据旋转特性来定
func search(nums []int, target int) int {
	n := len(nums)
	if n <= 0 {
		return -1
	}
	left := 0
	right := n - 1 // 包括右侧
	for left <= right {  // ==表示最后一个
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {  // 左..中
			if target >= nums[left] && target < nums[mid]  {  // 左..target..中
				right = mid - 1  // =mid已经返回
			} else {   // 两种可能target..左..右、左..右..target，都在右边
				left = mid + 1  // =mid已经返回
			}
		} else {  // 旋转点在左侧，7..4..8
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

