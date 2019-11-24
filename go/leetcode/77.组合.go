/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode-cn.com/problems/combinations/description/
 *
 * algorithms
 * Medium (68.47%)
 * Likes:    125
 * Dislikes: 0
 * Total Accepted:    12.8K
 * Total Submissions: 18.6K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 * 
 * 示例:
 * 
 * 输入: n = 4, k = 2
 * 输出:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 * 
 */

// @lc code=start
func combine(n int, k int) [][]int {
	out := []int{}
	res := [][]int{}
	dfs(n, k, &out, &res)
	return res
}

func dfs(n int, k int, out *[]int, res *[][]int) {
	if k <= 0 {
		*res = append(*res, append([]int{}, (*out)...))
		return
	}
	for i := n; i > 0; i-- {
		*out = append(*out, i)
		dfs(i-1, k-1, out, res)
		*out = (*out)[:len(*out)-1]
	}
}
// @lc code=end

