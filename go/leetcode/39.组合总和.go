/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode-cn.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (64.93%)
 * Likes:    284
 * Dislikes: 0
 * Total Accepted:    21.6K
 * Total Submissions: 33K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 * 
 * candidates 中的数字可以无限制重复被选取。
 * 
 * 说明：
 * 
 * 
 * 所有数字（包括 target）都是正整数。
 * 解集不能包含重复的组合。 
 * 
 * 
 * 示例 1:
 * 
 * 输入: candidates = [2,3,6,7], target = 7,
 * 所求解集为:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 * 
 * 
 * 示例 2:
 * 
 * 输入: candidates = [2,3,5], target = 8,
 * 所求解集为:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 * 
 */

// 求所有解的组合一般可以考虑使用递归实现，先往解集合中添加某个元素，递归后如果超过限制则回退，恰好符合则将解加入结果集，若未到达限制则继续尝试添加。
// 由于本题已经限定无重复元素，且元素可以重复选取，那么就是每次先递归同一个元素
func helper(arr []int, target int, start int, out *[]int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*res = append(*res, append([]int{}, (*out)...))  // 复制值
		return 
	}
	for i := start; i < len(arr); i++ {
		*out = append(*out, arr[i])
		helper(arr, target-arr[i], i, out, res)
		*out = (*out)[:len(*out)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	out := []int{}
	res := [][]int{}
	helper(candidates, target, 0, &out, &res)
	return res
}

