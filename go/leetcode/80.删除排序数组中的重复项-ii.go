/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/description/
 *
 * algorithms
 * Medium (51.91%)
 * Likes:    147
 * Dislikes: 0
 * Total Accepted:    24.7K
 * Total Submissions: 46.7K
 * Testcase Example:  '[1,1,1,2,2,3]'
 *
 * 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。
 * 
 * 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
 * 
 * 示例 1:
 * 
 * 给定 nums = [1,1,1,2,2,3],
 * 
 * 函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。
 * 
 * 你不需要考虑数组中超出新长度后面的元素。
 * 
 * 示例 2:
 * 
 * 给定 nums = [0,0,1,1,1,1,2,3,3],
 * 
 * 函数应返回新长度 length = 7, 并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。
 * 
 * 你不需要考虑数组中超出新长度后面的元素。
 * 
 * 
 * 说明:
 * 
 * 为什么返回数值是整数，但输出的答案是数组呢?
 * 
 * 请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
 * 
 * 你可以想象内部操作如下:
 * 
 * // nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
 * int len = removeDuplicates(nums);
 * 
 * // 在函数里修改输入数组对于调用者是可见的。
 * // 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
 * for (int i = 0; i < len; i++) {
 * print(nums[i]);
 * }
 * 
 */

// package leetcode
// 数组本身是拍好序的，现在限制最多两个重复，那么直接比较i和i-2，可以用双指针来实现
// @lc code=start
func removeDuplicates(nums []int) int {
	// 直接对比i和i-2的大小，如果==，则需要删掉i
	i := 0
	for j := 0; j < len(nums); j++ {
		if i < 2 || nums[j] > nums[i-2] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
// @lc code=end

