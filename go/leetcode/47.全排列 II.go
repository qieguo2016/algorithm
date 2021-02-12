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
	permuteDFS(&nums, &visited, &out, &res)
	return res
}

// 深度优先递归，通过访问标记数组去重
func permuteDFS(nums *[]int, visited *[]bool, out *[]int, res *[][]int) {
	if len(*out) == len(*nums) {
		*res = append(*res, append([]int{}, (*out)...))  // 复制值
		return 
	}
	for i := 0; i < len(*nums); i++ {
		if (*visited)[i] {
			continue
		}
		if (i > 0 && (*nums)[i] == (*nums)[i-1] && !(*visited)[i-1]) {
			// 用多叉树来描述解的空间：
			//    a0    a1(跳过)   b
			//   a1 b           a0 a1(跳过)
			//  b    a1        a1  
			// 横向第二个相同元素直接跳过，因为前面已经走过一次了，或者说a0/a1这两颗子树完全一样
			continue
		}
		(*visited)[i] = true
		*out = append(*out, (*nums)[i])
		permuteDFS(nums, visited, out, res)
		*out = (*out)[:len(*out)-1]
		(*visited)[i] = false
	}
}
// @lc code=end