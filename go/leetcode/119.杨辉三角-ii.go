/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 *
 * https://leetcode-cn.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (65.40%)
 * Likes:    293
 * Dislikes: 0
 * Total Accepted:    118.3K
 * Total Submissions: 180.9K
 * Testcase Example:  '3'
 *
 * 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 3
 * 输出: [1,3,3,1]
 *
 *
 * 进阶：
 *
 * 你可以优化你的算法到 O(k) 空间复杂度吗？
 *
 */

// @lc code=start
func getRow(rowIndex int) []int {
	//        1
	//       1 1
	//      1 2 1
	//     1 3 3 1
	// 	  1 4 6 4 1
	// var pre, cur []int
	// for i := 0; i <= rowIndex; i++ {  // 每次构造一行数据
	// 	cur = make([]int, i+1)
	// 	cur[0], cur[i] = 1, 1
	// 	for j := 1; j < i; j++ {
	// 		cur[j] = pre[j-1] + pre[j]
	// 	}
	// 	pre = cur
	// }

	// 优化空间，由cur[j] = pre[j-1] + pre[j]可知每次只和上一次结果有关
	// 这里如果使用递归的模式，就可以去掉pre数组
	// 每个数字是由它左肩上的数字加和得到，而递推是等于行内相加，所以不可以递推
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}
	return row
}

// @lc code=end

