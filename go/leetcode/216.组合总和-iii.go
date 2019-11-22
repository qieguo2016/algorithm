/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 *
 * https://leetcode-cn.com/problems/combination-sum-iii/description/
 *
 * algorithms
 * Medium (69.01%)
 * Likes:    71
 * Dislikes: 0
 * Total Accepted:    10.2K
 * Total Submissions: 14.8K
 * Testcase Example:  '3\n7'
 *
 * 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
 * 
 * 说明：
 * 
 * 
 * 所有数字都是正整数。
 * 解集不能包含重复的组合。 
 * 
 * 
 * 示例 1:
 * 
 * 输入: k = 3, n = 7
 * 输出: [[1,2,4]]
 * 
 * 
 * 示例 2:
 * 
 * 输入: k = 3, n = 9
 * 输出: [[1,2,6], [1,3,5], [2,3,4]]
 * 
 * 
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	out := []int{}
	res := [][]int{}
	dfs(k, n, 1, 10, &out, &res)
	return res
}

func dfs(k int, n int, start int, end int, out *[]int, res *[][]int) {
	if k == 0 && n == 0 {
		*res = append(*res, append([]int{}, (*out)...))
		return
	}
	if k < 0 || n < 0 {
		return
	}
	for i := start; i < end; i++ {
		*out = append(*out, i)
		dfs(k-1, n-i, i+1, end, out, res)
		*out = (*out)[:len(*out)-1]
	}
}
// @lc code=end

