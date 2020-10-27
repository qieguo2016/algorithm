/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

import (
	"sort"
)

// subsetsWithDup 首先排序，然后存上一个数字和当次遍历前的结果集s，当发现已经和上个数字相同的时候，则只处理s以外的结果
func subsetsWithDup(nums []int) [][]int {
  	ret := [][]int{}
	ret = append(ret, []int{})
	if len(nums) == 0 {
		return ret
	}
	sort.Ints(nums)
	last := nums[0]
	size := 1
	for i := 0; i < len(nums); i++ {
		if last != nums[i] {
			last = nums[i]
			size = len(ret)
		}
		newSize := len(ret)
		for j := newSize - size; j < newSize; j++ {
			sub := ret[j]
			sub = append(sub, nums[i])
			tmp := append([]int{}, sub...)
			ret = append(ret, tmp)
		}
	}
	return ret
}

