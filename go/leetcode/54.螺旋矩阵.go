/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode-cn.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (37.01%)
 * Likes:    215
 * Dislikes: 0
 * Total Accepted:    25.1K
 * Total Submissions: 67.6K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
 * 
 * 示例 1:
 * 
 * 输入:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * 输出: [1,2,3,6,9,8,7,4,5]
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * [
 * ⁠ [1, 2, 3, 4],
 * ⁠ [5, 6, 7, 8],
 * ⁠ [9,10,11,12]
 * ]
 * 输出: [1,2,3,4,8,12,11,10,9,5,6,7]
 * 
 * 
 */
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m <= 0 {
		return []int{}
	}
	n := len(matrix[0])
	if n <= 0 {
		return []int{}
	}
	left, right, up, down := 0, n-1, 0, m-1
	ret := make([]int, 0)
	for {
		// up -> right
		for j := left; j <= right;j++ {
			ret = append(ret, matrix[up][j])
		}
		up++
		if up > down {
			break
		}
		// right -> down
		for i := up; i <= down; i++  {
			ret = append(ret, matrix[i][right])
		}
		right--
		if left > right {
			break
		}
		// down -> left
		for j := right; j >= left; j-- {
			ret = append(ret, matrix[down][j])
		}
		down--
		if up > down {
			break
		}
		// left -> up
		for i := down; i >= up; i-- {
			ret = append(ret, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return ret
}

