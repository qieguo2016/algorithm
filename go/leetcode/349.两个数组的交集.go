/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (65.69%)
 * Likes:    128
 * Dislikes: 0
 * Total Accepted:    33.9K
 * Total Submissions: 51.2K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 * 
 * 示例 1:
 * 
 * 输入: nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出: [2]
 * 
 * 
 * 示例 2:
 * 
 * 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出: [9,4]
 * 
 * 说明:
 * 
 * 
 * 输出结果中的每个元素一定是唯一的。
 * 我们可以不考虑输出结果的顺序。
 * 
 * 
 */
func intersection(nums1 []int, nums2 []int) []int {
	m1 := map[int]bool{}
	for _, el := range nums1 {
		m1[el] = true
	}
	m2 := map[int]bool{}
	for _, el := range nums2 {
		if m1[el] {
			m2[el] = true
		}
	}
	ret := []int{}
	for k, _ := range m2 {
		ret = append(ret, k)
	}
	return ret
}

