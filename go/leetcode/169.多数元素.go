/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 *
 * https://leetcode-cn.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (64.84%)
 * Likes:    778
 * Dislikes: 0
 * Total Accepted:    228.4K
 * Total Submissions: 352.3K
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
 * 
 * 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: [3,2,3]
 * 输出: 3
 * 
 * 示例 2:
 * 
 * 输入: [2,2,1,1,1,2,2]
 * 输出: 2
 * 
 * 
 */

// @lc code=start
// hash记录
func majorityElement(nums []int) int {
	m := map[int]int{}
	ret := 0
	max := 0
	for _, n := range nums {
		v, ok := m[n];
		if ok {
			v = v+1
		} else {
			v = 1
		}
		m[n] = v
		if v > max {
			max = v
			ret = n
		}
	}
	return ret
}
// @lc code=end

