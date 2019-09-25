/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (70.38%)
 * Likes:    385
 * Dislikes: 0
 * Total Accepted:    24.1K
 * Total Submissions: 34.2K
 * Testcase Example:  '3'
 *
 * 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
 * 
 * 例如，给出 n = 3，生成结果为：
 * 
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 * 
 * 
 */

func helper(left int, right int, out string, res *[]string) {
	if left > right {
		return
	}
	if left == 0 && right == 0 {
		*res = append(*res, out) // 复制值
		return 
	}
	if left > 0 {
		helper(left-1, right, out+"(", res)
	}
	if right > 0 {
		helper(left, right-1, out+")", res)
	}
}

func generateParenthesis(n int) []string {
	res := []string{}
	helper(n, n, "", &res)
	return res
}

