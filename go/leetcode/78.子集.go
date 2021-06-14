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
	s := &solution{
		res:  [][]int{},
		path: []int{},
	}
	s.call(nums, 0)
	return s.res
}

type solution struct {
	path []int
	res  [][]int
}

func (s *solution) call(nums []int, start int) {
	if start > len(nums) {
		return
	}

	s.res = append(s.res, append([]int{}, s.path...))

	for i := start; i < len(nums); i++ {
		s.path = append(s.path, nums[i])
		s.call(nums, i+1)
		s.path = s.path[:len(s.path)-1]
	}
}
