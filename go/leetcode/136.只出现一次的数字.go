/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 *
 * https://leetcode-cn.com/problems/single-number/description/
 *
 * algorithms
 * Easy (63.24%)
 * Likes:    924
 * Dislikes: 0
 * Total Accepted:    121.3K
 * Total Submissions: 188.9K
 * Testcase Example:  '[2,2,1]'
 *
 * 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 * 
 * 说明：
 * 
 * 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 * 
 * 示例 1:
 * 
 * 输入: [2,2,1]
 * 输出: 1
 * 
 * 
 * 示例 2:
 * 
 * 输入: [4,1,2,1,2]
 * 输出: 4
 * 
 */

// @lc code=start
// 异或：相同的位置置0，不同的位置1
// 因此，对数组所有元素进行异或，出现2次的数字的所有位被抵消了，只剩出现1次的数字
func singleNumber(nums []int) int {
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		ret ^= nums[i]
	}
	return ret
}
// @lc code=end

