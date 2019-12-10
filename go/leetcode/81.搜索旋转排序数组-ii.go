/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Medium (34.13%)
 * Likes:    64
 * Dislikes: 0
 * Total Accepted:    10.7K
 * Total Submissions: 31.4K
 * Testcase Example:  '[2,5,6,0,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 *
 * ( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
 *
 * 编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
 *
 * 示例 1:
 *
 * 输入: nums = [2,5,6,0,0,1,2], target = 0
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: nums = [2,5,6,0,0,1,2], target = 3
 * 输出: false
 *
 * 进阶:
 *
 *
 * 这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
 * 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
 *
 *
 */

/*
 * 二分查找加上判断旋转点
 * 可以判断左中右三个点的关系判断旋转点，左小于中一直递增，左大于中存在旋转点
 */
// package leetcode

// @lc code=start
func search(nums []int, target int) bool {
	length := len(nums)
	if length <= 0 {
		return false
	}
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[left] == target || nums[right] == target || nums[mid] == target {
			return true
		}
		// 当中点==左边界的时候，有可能是[1,3,1,1]的情况，右边界左移一位一直到可以继续位置
		if nums[mid] == nums[left] {
			right--
			continue
		}
		if nums[mid] > nums[left] { // 旋转点在右侧
			if target > nums[left] && target < nums[mid] { // 目标点在左侧
				right = mid - 1
			} else {
				left = mid + 1
			}
			continue
		}
		// 旋转点在左侧
		if target > nums[mid] && target < nums[right] { // 目标点在右侧
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// @lc code=end
