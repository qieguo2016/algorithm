/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (38.17%)
 * Likes:    192
 * Dislikes: 0
 * Total Accepted:    30K
 * Total Submissions: 77.8K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 *
 * 示例 1:
 *
 * 输入: [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2:
 *
 * 输入: [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 */

import (
	"sort"
)

type SortArray [][]int

func (a SortArray) Len() int {
	return len(a)
}
func (a SortArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortArray) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}

func merge(intervals [][]int) [][]int {
	sort.Sort(SortArray(intervals))
	ret := [][]int{}
	tmp := []int{}
	for i := 0; i < len(intervals); i++ {
		if len(tmp) <= 0 {
			tmp = intervals[i]
		}
		if i >= len(intervals) - 1 {
			ret = append(ret, tmp)
			break
		}
		if tmp[1] >= intervals[i+1][0] { // 有重叠
			if tmp[1] < intervals[i+1][1] {
				tmp[1] = intervals[i+1][1]
			}
		} else {
			ret = append(ret, tmp)
			tmp = []int{}
		}
	}
	return ret
}
