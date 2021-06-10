/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 *
 * https://leetcode-cn.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (36.29%)
 * Likes:    269
 * Dislikes: 0
 * Total Accepted:    23.4K
 * Total Submissions: 64.1K
 * Testcase Example:  '[1,2,0]'
 *
 * 给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
 *
 * 示例 1:
 *
 * 输入: [1,2,0]
 * 输出: 3
 *
 *
 * 示例 2:
 *
 * 输入: [3,4,-1,1]
 * 输出: 2
 *
 *
 * 示例 3:
 *
 * 输入: [7,8,9,11,12]
 * 输出: 1
 *
 *
 * 说明:
 *
 * 你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。
 *
 */

// 原地交换排序，两次遍历
// 交换后的数组要符合[1,2,3,4... ]的规则，第一个不符合这样规则的位置，就是结果
func firstMissingPositive(nums []int) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		for nums[i] > 0 && nums[i] < size && nums[i] != nums[nums[i]-1] { // 1应该在0位
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < size; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return size + 1
}

