/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N皇后
 *
 * https://leetcode-cn.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (64.97%)
 * Likes:    232
 * Dislikes: 0
 * Total Accepted:    14.9K
 * Total Submissions: 22.7K
 * Testcase Example:  '4'
 *
 * n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 * 
 * 
 * 
 * 上图为 8 皇后问题的一种解法。
 * 
 * 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
 * 
 * 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 * 
 * 示例:
 * 
 * 输入: 4
 * 输出: [
 * ⁠[".Q..",  // 解法 1
 * ⁠ "...Q",
 * ⁠ "Q...",
 * ⁠ "..Q."],
 * 
 * ⁠["..Q.",  // 解法 2
 * ⁠ "Q...",
 * ⁠ "...Q",
 * ⁠ ".Q.."]
 * ]
 * 解释: 4 皇后问题存在两个不同的解法。
 * 
 * 
 */

// 几点性质
// 1. 必定是每行一个皇后
// 2. 必定是每列一个皇后
func solveNQueens(n int) [][]string {
	rect := make([][]string, n)  // 棋盘
	for i := 0; i < n; i++ {
		rect[i] = make([]string, n)
		for j := 0; j < n; j++ {
			rect[i][j] = "."
		}
	}
	res := [][]string{}
	dfs(&rect, 0, &res)
	return res
}

func dfs(rect *[][]string, row int, res *[][]string) {
	if row >= len(*rect) {
		// 输出一个解
		ret := []string{}
		for _, row := range (*rect) {
			str := ""
			for _, col := range row {
				str += string(col)
			}
			ret = append(ret, str)
		}
		*res = append(*res, ret)
		return 
	}
	for i := 0; i < len(*rect); i++ {  // 试探各个列
		if isValid(rect, row, i) {
			(*rect)[row][i] = "Q"
			dfs(rect, row+1, res)
			(*rect)[row][i] = "."
		}
	}
}

func isValid(rect *[][]string, row int, col int) bool {
	// 因为每次都是递增row和col，所以之后的row和col不需要检查
	// 不需要检查行，以为每行只有一次放的机会
	// 检测列
	for i := 0; i < row; i++ {
		if (*rect)[i][col] == "Q" {
			return false
		}
	}
	// 左上对角线检查
	for i, j := row - 1, col - 1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if (*rect)[i][j] == "Q" {
			return false
		}
	}
	// 右上对角线检查
	for i, j := row - 1, col + 1; i >= 0 && j < len(*rect); i, j = i-1, j+1 {
		if (*rect)[i][j] == "Q" {
			return false
		}
	}
	return true
}

