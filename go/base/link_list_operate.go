package base

import (
	"container/heap"
	"errors"
)

// ReverseLinkList 原地反转链表
// 新建一个返回链表，next指向原链表头
// 每次循环：原链头所指向的元素移动到返回链头，并指向上次循环的返回链头
func ReverseLinkList(head *LinkListNode) *LinkListNode {
	dummy := &LinkListNode{}
	dummy.next = head
	for head != nil && head.next != nil {
		n := head.next
		head.next = head.next.next
		n.next = dummy.next
		dummy.next = n
	}
	return dummy.next
}

// IsLinkListLoop 是否存在闭环，两种实现
// 1 增加一个外部存储(hash map)标记已遍历的节点看是否有重复
// 2 使用两个指针，采用不同的步进长度，如果两者相遇则表示有环路
func IsLinkListLoop(head *LinkListNode) bool {
	c1 := head
	if c1 == nil {
		return false
	}
	c2 := c1.next
	for ; c1 != nil; c1 = c1.next { // 步进长度为1
		if c1.IsEqual(c2) {
			return true
		}
		// 只有一个元素
		if c2 == nil {
			return false
		}
		// 只有两个元素
		if c2.next == nil {
			return false
		}
		c2 = c2.next.next // 步进长度为2
	}
	return false
}

// MergeSortedLinkList 合并有序链表
// 1: 归并排序，开一个新的链表，每次将小的放进去
// 2: 原地归并，使用其中一条作为基线进行归并，找到两者的小值之后，再在基线上移动游标寻找插入位置
func MergeSortedLinkList(l1 *LinkList, l2 *LinkList) *LinkList {
	c1 := l1.head
	c2 := l2.head

	// // 普通归并
	// head := &LinkListNode{}
	// c := head
	// for c1 != nil && c2 != nil {
	// 	if c1.value < c2.value {
	// 		c.next = &LinkListNode{value: c1.value}
	// 		c = c.next
	// 		c1 = c1.next
	// 	} else {
	// 		c.next = &LinkListNode{value: c2.value}
	// 		c = c.next
	// 		c2 = c2.next
	// 	}
	// }
	// if c1 != nil {
	// 	c.next = c1
	// }
	// if c2 != nil {
	// 	c.next = c2
	// }
	// return &LinkList{length: l1.length + l2.length, head: head.next}

	// 空间优化解法，基准是l1
	pos := l1.head
	for c1 != nil && c2 != nil {
		if c1.value < c2.value {
			if pos == c1 {
				continue
			}
			pos.next = c1
			pos = pos.next
			c1 = c1.next
			continue
		}
		if pos == l1.head {
			l1.head = c2
			pos = l1.head
			c2 = c2.next
			continue
		}
		pos.next = c2
		pos = pos.next
		c2 = c2.next
	}
	if c1 != nil {
		pos.next = c1
	}
	if c2 != nil {
		pos.next = c2
	}

	return &LinkList{length: l1.length + l2.length, head: l1.head}
}

// GetRevKthFromLinkList 获取链表倒数第k个数
func GetRevKthFromLinkList(head *LinkListNode, k int) (int, error) {
	if head == nil || k < 0 {
		return 0, errors.New("not_found")
	}
	n := 0
	ret := head.value
	for head.next != nil && n < k {
		head = head.next
		n++
	}
	for head.next != nil {
		head = head.next
		ret = head.value
	}
	if n == k {
		return ret, nil
	}
	return 0, errors.New("not_found")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

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
// 1. 采用小根堆存储每个链表的队头，每次从小根堆头部取节点加入返回链表中，并将该节点的Next放回小根堆中
// 2. 也可以使用分治法，将链表按照两两划分合并，然后递归合并到只剩一个链表
func MergeKLists(lists []*ListNode) *ListNode {
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
