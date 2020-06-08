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

// 树的根节点是左括号，层数是3层，求路径集合，终止条件：左右数递减到0，或者右节点比左节点多
func helper(left int, right int, out string, res *[]string) {
	if left > right {  // 左剩余比右多，也就是右节点多了，没法匹配
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

// 组合题目一般用递归，组合一般可以看成是树结构，用dfs、bfs配合规则求解
func generateParenthesis(n int) []string {
	res := []string{}
	helper(n, n, "", &res)
	return res
}

