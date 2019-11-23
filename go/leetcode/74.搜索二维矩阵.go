/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 *
 * https://leetcode-cn.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (35.89%)
 * Likes:    85
 * Dislikes: 0
 * Total Accepted:    17.3K
 * Total Submissions: 48.2K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,50]]\n3'
 *
 * 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
 * 
 * 
 * 每行中的整数从左到右按升序排列。
 * 每行的第一个整数大于前一行的最后一个整数。
 * 
 * 
 * 示例 1:
 * 
 * 输入:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 3
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 13
 * 输出: false
 * 
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		return false
	}
	ml, mr := 0, len(matrix)-1  // left/right
	for ml < mr {  // 0,2 | 
		mm := ml + (mr-ml)/2  // mid: 1, 1
		if matrix[mm][0] == target {
			return true
		}
		if matrix[mm][0] > target {
			mr = mm - 1
			continue
		}
		if ml < mm {
			ml = mm  // 1
			continue
		}
		// 最后两行，ml=mm，判断是否在mr
		if matrix[mr][0] == target {
			return true  // 小剪枝
		}
		if matrix[mr][0] > target {
			mr = mm  // 1
			continue
		}
		ml = mm + 1
	}
	// find in matrix[ml]
	row := matrix[ml]
	left, right := 0, len(row)-1
	for left <= right {
		mid := left + (right-left)/2
		if row[mid] == target {
			return true
		}
		if row[mid] < target {
			left = mid + 1
			continue
		}
		right = mid - 1
	}
	return false
}
// @lc code=end

