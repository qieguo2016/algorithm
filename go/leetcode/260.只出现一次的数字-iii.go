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
// 先将数组整体异或得到两个不同数字的异或结果，因为这两个数字不同，所以异或结果必然不为0
// 异或结果中1的位就是两者不同的位，也就是某个数为1的位，找一个这样的位然后将其他位置为0，得到一个flag
// 遍历数组，用这个flag去与，如果不为0则表示这个位是1，为0则表示这个位是0，这样就将不同的两个数分在两个组里面了
// 对这两个组分别组内异或，就能得到两个不同的数字
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

