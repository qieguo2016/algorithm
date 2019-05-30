/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import (
	"container/heap"
)

type srHeap []*ListNode

func (h *srHeap) Less(i, j int) bool {
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
	*h = append(*h, v.(*ListNode))
}

// MergeKLists 合并k个有序链表
func mergeKLists(lists []*ListNode) *ListNode {
	heads := make(srHeap, 0)
	for _, h := range lists {
		if h != nil {
			heads = append(heads, h)
		}
	}
	heap.Init(&heads)
	dummy := &ListNode{}
	c := dummy
	for len(heads) > 0 {
		c.Next = heap.Pop(&heads).(*ListNode)
		c = c.Next
		if c.Next != nil {
			heap.Push(&heads, c.Next)
		}
	}
	return dummy.Next
}
