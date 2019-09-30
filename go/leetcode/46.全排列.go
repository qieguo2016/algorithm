/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (71.53%)
 * Likes:    365
 * Dislikes: 0
 * Total Accepted:    43.4K
 * Total Submissions: 60.7K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个没有重复数字的序列，返回其所有可能的全排列。
 * 
 * 示例:
 * 
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 * 
 */


func permute(nums []int) [][]int {
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
		(*visited)[i] = true
		*out = append(*out, (*nums)[i])
		permuteDFS(nums, level+1, visited, out, res)
		*out = (*out)[:len(*out)-1]
		(*visited)[i] = false
	}
}

