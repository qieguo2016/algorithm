/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode-cn.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (74.51%)
 * Likes:    120
 * Dislikes: 0
 * Total Accepted:    16.4K
 * Total Submissions: 21.9K
 * Testcase Example:  '3'
 *
 * 给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
 *
 * 示例:
 *
 * 输入: 3
 * 输出:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 8, 9, 4 ],
 * ⁠[ 7, 6, 5 ]
 * ]
 *
 */

package leetcode

// @lc code=start
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	left, right, up, down := 0, n-1, 0, n-1
	i := 1
	for {
		for j := left; j <= right; j++ {
			res[up][j] = i
			i++
		}
		up++
		if up > down {
			break
		}
		for j := up; j <= down; j++ {
			res[j][right] = i
			i++
		}
		right--
		if left > right {
			break
		}
		for j := right; j >= left; j-- {
			res[down][j] = i
			i++
		}
		down--
		if up > down {
			break
		}
		for j := down; j >= up; j-- {
			res[j][left] = i
			i++
		}
		left++
		if left > right {
			break
		}
	}
	return res
}

// @lc code=end
