/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 *
 * https://leetcode-cn.com/problems/permutation-sequence/description/
 *
 * algorithms
 * Medium (46.71%)
 * Likes:    123
 * Dislikes: 0
 * Total Accepted:    14.3K
 * Total Submissions: 30.4K
 * Testcase Example:  '3\n3'
 *
 * 给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
 * 
 * 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
 * 
 * 
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 * 
 * 
 * 给定 n 和 k，返回第 k 个排列。
 * 
 * 说明：
 * 
 * 
 * 给定 n 的范围是 [1, 9]。
 * 给定 k 的范围是[1,  n!]。
 * 
 * 
 * 示例 1:
 * 
 * 输入: n = 3, k = 3
 * 输出: "213"
 * 
 * 
 * 示例 2:
 * 
 * 输入: n = 4, k = 9
 * 输出: "2314"
 * 
 * 
 */
func getPermutation(n int, k int) string {
	// 1234 1243 1324 1342 1423 1432
	// 第n位固定、其他位的取值有(n-1)!种
	nums := []byte("123456789")
	c := make([]int, n)
	c[0] = 1
	// 预先算好阶乘
	for i := 1; i < n; i++ {
		c[i] = c[i-1] * i
	}
	ret := ""
	k--  // 序号从0开始
	for i := n-1; i >= 0; i-- {
		j := k / c[i]  // 8/6  2/2
		ret += string(nums[j])  // 2 3 
		if j + 1 < len(nums) {  // 134
			nums = append(nums[:j], nums[j+1: ]...)
		} else {
			nums = nums[:j]
		}
		k %= c[i]  // 2
	}
	return ret
}

