/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// subsets 从空集开始，遍历时每次对结果集合内的每个结果继续叠加
func subsets(nums []int) [][]int {
	ret := [][]int{}
	ret = append(ret, []int{})
	for i := 0; i < len(nums); i++ {
		for _, sub := range ret {
			sub = append(sub, nums[i])
			tmp := append([]int{}, sub...)
			ret = append(ret, tmp)
		}
	}
	return ret
}

