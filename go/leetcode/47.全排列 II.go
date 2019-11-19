/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (53.86%)
 * Likes:    164
 * Dislikes: 0
 * Total Accepted:    23.5K
 * Total Submissions: 43.3K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列，返回所有不重复的全排列。
 * 
 * 示例:
 * 
 * 输入: [1,1,2]
 * 输出:
 * [
 * ⁠ [1,1,2],
 * ⁠ [1,2,1],
 * ⁠ [2,1,1]
 * ]
 * 
 */

// @lc code=start
import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	visited := make([]bool, len(nums))
	out := []int{}
	res := [][]int{}
	permuteDFS(&nums, 0, &visited, &out, &res)
	return res
}

// 深度优先递归，通过访问标记数组去重
func permuteDFS(nums *[]int, level int, visited *[]bool, out *[]int, res *[][]int) {
	if level == len(*nums) {
		*res = append(*res, append([]int{}, (*out)...))  // 复制值
		return 
	}
	for i := 0; i < len(*nums); i++ {
		if (*visited)[i] {
			continue
		}
		if (i > 0 && (*nums)[i] == (*nums)[i-1] && !(*visited)[i-1]) {
			continue
		}
		(*visited)[i] = true
		*out = append(*out, (*nums)[i])
		permuteDFS(nums, level+1, visited, out, res)
		*out = (*out)[:len(*out)-1]
		(*visited)[i] = false
	}
}
// @lc code=end