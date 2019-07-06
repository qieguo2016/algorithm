/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 *
 * https://leetcode-cn.com/problems/top-k-frequent-words/description/
 *
 * algorithms
 * Medium (39.49%)
 * Likes:    36
 * Dislikes: 0
 * Total Accepted:    1.9K
 * Total Submissions: 4.9K
 * Testcase Example:  '["i", "love", "leetcode", "i", "love", "coding"]\n2'
 *
 * 给一非空的单词列表，返回前 k 个出现次数最多的单词。
 *
 * 返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。
 *
 * 示例 1：
 *
 *
 * 输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
 * 输出: ["i", "love"]
 * 解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
 * ⁠   注意，按字母顺序 "i" 在 "love" 之前。
 *
 *
 *
 *
 * 示例 2：
 *
 *
 * 输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"],
 * k = 4
 * 输出: ["the", "is", "sunny", "day"]
 * 解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
 * ⁠   出现次数依次为 4, 3, 2 和 1 次。
 *
 *
 *
 *
 * 注意：
 *
 *
 * 假定 k 总为有效值， 1 ≤ k ≤ 集合元素数。
 * 输入的单词均由小写字母组成。
 *
 *
 *
 *
 * 扩展练习：
 *
 *
 * 尝试以 O(n log k) 时间复杂度和 O(n) 空间复杂度解决。
 *
 *
 */

// package leetcode

import (
	"container/heap"
)

type HeapNode struct {
	Data string
	Val  int
}

type srHeap []*HeapNode

func (h *srHeap) Less(i, j int) bool {
	// 相同Val的节点按字典序逆序
	if (*h)[i].Val == (*h)[j].Val {
		return (*h)[i].Data > (*h)[j].Data
	}
	return (*h)[i].Val < (*h)[j].Val
}

func (h *srHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *srHeap) Len() int {
	return len(*h)
}

func (h *srHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *srHeap) Push(v interface{}) {
	*h = append(*h, v.(*HeapNode))
}

func topKFrequent(words []string, k int) []string {
	if len(words) <= 1 {
		return words
	}
	wc := map[string]int{}
	// hashmap存储word、count
	for _, w := range words {
		_, ok := wc[w]
		if !ok {
			wc[w] = 0
		}
		wc[w]++
	}
	hp := make(srHeap, 0)
	heap.Init(&hp)
	for w, c := range wc {
		// 超过k，则比较顶点值，小于顶点值跳过，等于顶点值且字典序小、大于顶点值加入堆
		if len(hp) >= k {
			if hp[0].Val == c && w > hp[0].Data {
				continue
			} else if hp[0].Val > c {
				continue
			}
			heap.Pop(&hp)
		}
		heap.Push(&hp, &HeapNode{Val: c, Data: w})
	}
	ret := []string{}
	for len(hp) > 0 {
		el := heap.Pop(&hp).(*HeapNode)
		ret = append(ret, el.Data)
	}
	rev := []string{}
	i := len(ret) - 1
	for i >= 0 {
		rev = append(rev, ret[i])
		i--
	}
	return rev
}
