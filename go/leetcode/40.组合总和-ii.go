/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode-cn.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (55.85%)
 * Likes:    109
 * Dislikes: 0
 * Total Accepted:    15.1K
 * Total Submissions: 27K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用一次。
 *
 * 说明：
 *
 *
 * 所有数字（包括目标数）都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1:
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 所求解集为:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 所求解集为:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 */

import (
	"sort"
)

type solution struct {
	out []int
	res [][]int
}

func (s *solution) helper(arr []int, target int, start int) {
	if target < 0 {
		return
	}
	if target == 0 {
		s.res = append(s.res, append([]int{}, (s.out)...)) // 复制值
		return
	}
	for i := start; i < len(arr); i++ {
		if i > start && arr[i] == arr[i-1] {
			continue
		}
		s.out = append(s.out, arr[i])
		s.helper(arr, target-arr[i], i+1)
		s.out = s.out[:len(s.out)-1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	s := &solution{
		out: []int{},
		res: [][]int{},
	}
	s.helper(candidates, target, 0)
	return s.res
}

