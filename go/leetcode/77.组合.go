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
	s := &solution{
		out: []int{},
		res: [][]int{},
	}
	s.call(n, k)
	return s.res
}

type solution struct {
	out []int
	res [][]int
}

func (s *solution) call(n int, k int) {
	if k <= 0 {
		s.res = append(s.res, append([]int{}, s.out...))
		return
	}
	for i := n; i >= k; i-- {
		s.out = append(s.out, i)
		s.call(i-1, k-1)
		s.out = s.out[:len(s.out)-1]
	}
}

// @lc code=end

