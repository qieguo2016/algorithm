/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 *
 * https://leetcode-cn.com/problems/single-number-ii/description/
 *
 * algorithms
 * Medium (64.83%)
 * Likes:    210
 * Dislikes: 0
 * Total Accepted:    15.8K
 * Total Submissions: 24.4K
 * Testcase Example:  '[2,2,3,2]'
 *
 * 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
 * 
 * 说明：
 * 
 * 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 * 
 * 示例 1:
 * 
 * 输入: [2,2,3,2]
 * 输出: 3
 * 
 * 
 * 示例 2:
 * 
 * 输入: [0,1,0,1,0,1,99]
 * 输出: 99
 * 
 */

// @lc code=start
func singleNumber(nums []int) int {
	ret := 0
	// 遍历32个位，分别统计每个位上1的个数，除掉特殊的那个数字，其他数字都出现了3次，所以个数必然是3的倍数
	// 只要将个数模3，就可以得到特殊数在这个位上是否为1了
	for i := 0; i < 64; i++ {  // 注意64位机器上int就是64位，如果写的32则在64位机器会失败
		sum := 0
		for _, num := range nums {
			sum += (num >> i) & 1 // 第i位是否为1
		}
		ret |= (sum % 3) << i  // sum%3是目标数在该位的值
	}
	return ret
}
// @lc code=end

