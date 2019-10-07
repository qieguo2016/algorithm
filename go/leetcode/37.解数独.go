/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 *
 * https://leetcode-cn.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (55.97%)
 * Likes:    242
 * Dislikes: 0
 * Total Accepted:    11.4K
 * Total Submissions: 20.1K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * 编写一个程序，通过已填充的空格来解决数独问题。
 * 
 * 一个数独的解法需遵循如下规则：
 * 
 * 
 * 数字 1-9 在每一行只能出现一次。
 * 数字 1-9 在每一列只能出现一次。
 * 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
 * 
 * 
 * 空白格用 '.' 表示。
 * 
 * 
 * 
 * 一个数独。
 * 
 * 
 * 
 * 答案被标成红色。
 * 
 * Note:
 * 
 * 
 * 给定的数独序列只包含数字 1-9 和字符 '.' 。
 * 你可以假设给定的数独只有唯一解。
 * 给定数独永远是 9x9 形式的。
 * 
 * 
 */

// @lc code=start
// 回溯法经典案例
const SIZE = 9
const empty byte = '.'
const one byte = '1'
// 检查某个数是否合法
func isValid(bp *[][]byte, row int, col int, val byte) bool {
	board := *bp
	// 行
	for j := 0; j < SIZE; j++ {
		if board[row][j] == val {
			return false
		}
	}
	// 列
	for i := 0; i < SIZE; i++ {
		if board[i][col] == val {
			return false
		}
	}
	// 块 2,5  [0-3) [3-6) 
	br := row/3*3
	bc := col/3*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[br+i][bc+j] == val {
				return false
			}
		}
	}
	return true
}
// 填入一个数之后检查横竖块是否符合，然后填下一个，如果下一个无解，则回退上一个数继续尝试
func helper(bp *[][]byte, i int, j int) bool {
	// 每次在行内移动，遍历各个列，到行尾（最后一列）之后再换行
	if i >= SIZE {
		return true
	}
	if j >= SIZE {
		return helper(bp, i+1, 0)
	}
	board := *bp
	if board[i][j] != empty {
		return helper(bp, i, j+1)
	}
	// 尝试填入1-9
	for k := 0 ; k < SIZE; k++ {
		b := one + byte(k)
		if !isValid(bp, i, j, b) {
			continue
		}
		board[i][j] = b  // 确定这个值，继续填下一个
		if helper(bp, i, j+1) {
			return true
		}
		board[i][j] = empty // 回退
	}
	return false
}

func solveSudoku(board [][]byte)  {
	helper(&board, 0, 0)
}
// @lc code=end

