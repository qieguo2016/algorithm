/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 *
 * https://leetcode-cn.com/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (75.15%)
 * Likes:    71
 * Dislikes: 0
 * Total Accepted:    9.6K
 * Total Submissions: 12.7K
 * Testcase Example:  '4'
 *
 * n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 * 
 * 
 * 
 * 上图为 8 皇后问题的一种解法。
 * 
 * 给定一个整数 n，返回 n 皇后不同的解决方案的数量。
 * 
 * 示例:
 * 
 * 输入: 4
 * 输出: 2
 * 解释: 4 皇后问题存在如下两个不同的解法。
 * [
 * [".Q..",  // 解法 1
 * "...Q",
 * "Q...",
 * "..Q."],
 * 
 * ["..Q.",  // 解法 2
 * "Q...",
 * "...Q",
 * ".Q.."]
 * ]
 * 
 * 
 */
func totalNQueens(n int) int {
	rect := make([][]int, n)  // 棋盘
	for i := 0; i < n; i++ {
		rect[i] = make([]int, n)
	}
	res := 0
	dfs(0, &rect, &res)
	return res
}

func dfs(row int, rect *[][]int, res *int)  {
	if row >= len(*rect) {
		*res++
		return 
	}
	for i := 0; i < len(*rect); i++ {
		if isValid(rect, row, i) {
			(*rect)[row][i] = 1
			dfs(row+1, rect, res)
			(*rect)[row][i] = 0
		}
	}
}

func isValid(rect *[][]int, row int, col int) bool {
	for i := 0; i < row; i++ {
		if (*rect)[i][col] == 1 {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if (*rect)[i][j] == 1 {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < len(*rect); i, j = i-1, j+1 {
		if (*rect)[i][j] == 1 {
			return false
		}
	}
	return true
}

