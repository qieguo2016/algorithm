/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 *
 * https://leetcode-cn.com/problems/single-number-iii/description/
 *
 * algorithms
 * Medium (67.85%)
 * Likes:    141
 * Dislikes: 0
 * Total Accepted:    10.9K
 * Total Submissions: 16.1K
 * Testcase Example:  '[1,2,1,3,2,5]'
 *
 * 给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。
 * 
 * 示例 :
 * 
 * 输入: [1,2,1,3,2,5]
 * 输出: [3,5]
 * 
 * 注意：
 * 
 * 
 * 结果输出的顺序并不重要，对于上面的例子， [5, 3] 也是正确答案。
 * 你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
 * 
 * 
 */

// @lc code=start
func singleNumber(nums []int) []int {
    diff := nums[0]
	for i := 1; i < len(nums); i++ {
		diff ^= nums[i]
	}
	diff &= -diff  // 保留从右往左第一个1，其他位设为0
	ret := make([]int, 2)
	for _, num := range nums {
		if num & diff == 0  {
			ret[0] ^= num
		} else {
			ret[1] ^= num
		}
	}
	return ret
}
// @lc code=end

