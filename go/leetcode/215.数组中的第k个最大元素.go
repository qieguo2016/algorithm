/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (59.34%)
 * Likes:    293
 * Dislikes: 0
 * Total Accepted:    58.7K
 * Total Submissions: 98.9K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 * 示例 1:
 *
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 *
 *
 * 示例 2:
 *
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 *
 * 说明:
 *
 * 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
 *
 */

// top k，第一反应用堆，构造一个大小为k的最小堆，判断堆顶元素
// package leetcode

// @lc code=start
import (
	// "container/heap"
	"math/rand"
)

// type myHeap []int

// func (h *myHeap) Less(i, j int) bool {
// 	return (*h)[i] < (*h)[j]
// }

// func (h *myHeap) Swap(i, j int) {
// 	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
// }

// func (h *myHeap) Len() int {
// 	return len(*h)
// }

// func (h *myHeap) Pop() (v interface{}) {
// 	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
// 	return
// }

// func (h *myHeap) Push(v interface{}) {
// 	*h = append(*h, v.(int))
// }

func findKthLargest(nums []int, k int) int {
	// h := make(myHeap, 0)
	// heap.Init(&h)
	// for _, cur := range nums {
	// 	if h.Len() >= k {
	// 		if cur <= h[0] {
	// 			continue
	// 		}
	// 		heap.Pop(&h)
	// 	}
	// 	heap.Push(&h, cur)
	// }
	// return h[0]
	s := &solution{nums: nums}
	return s.call(len(nums)-k, 0, len(nums)-1)
}

type solution struct {
	nums []int
}

// 用快排的分组思想，如果分组的界限是k，就可以返回了，否则继续2分。
func (s *solution) call(k, left, right int) int {
	idx := s.partition(left, right)
	if k == idx {
		return s.nums[k]
	}
	if k > idx {
		return s.call(k, idx+1, right)
	}
	return s.call(k, left, idx-1)
}

// 左闭右闭
func (s *solution) partition(i, j int) int {
	if i >= j {
		return i
	}
	idx := rand.Intn(j-i) + i
	s.nums[idx], s.nums[i] = s.nums[i], s.nums[idx]
	d := s.nums[i]
	for i < j {
		for i < j && s.nums[j] >= d {
			j--
		}
		s.nums[i] = s.nums[j]
		for i < j && s.nums[i] < d {
			i++
		}
		s.nums[j] = s.nums[i]
	}
	s.nums[i] = d
	return i
}

// @lc code=end
