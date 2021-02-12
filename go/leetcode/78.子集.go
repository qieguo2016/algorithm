/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// subsets 从空集开始，遍历时每次对结果集合内的每个结果继续叠加
// func subsets(nums []int) [][]int {
// 	ret := [][]int{}
// 	ret = append(ret, []int{})
// 	for i := 0; i < len(nums); i++ {
// 		for _, sub := range ret {
// 			sub = append(sub, nums[i])
// 			tmp := append([]int{}, sub...)
// 			ret = append(ret, tmp)
// 		}
// 	}
// 	return ret
// }

// 用回溯法，组合模式来做
func subsets(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	helper(nums, &result, &path, 0)
	return result
}

func helper(nums []int, result *[][]int, path *[]int, start int) {
	if start > len(nums) {
		return
	}
	
	*result = append(*result, append([]int{}, (*path)...))

	for i := start; i < len(nums); i++ {
		*path = append(*path, nums[i])
		helper(nums, result, path, i+1)
		*path = (*path)[:len(*path)-1]
	}
}
