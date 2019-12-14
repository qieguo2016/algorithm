package list

import (
	"container/heap"
	"errors"
)

// 链表操作类：1.调整链表顺序, 2.两个链表相互作用，3）判断链表性质
// 常见问题：
// a)单向链表操作基础是前一个指针，所以一般从head开始循环cur.Next，记得判断cur.Next!=nil
// b)操作数从1开始，都是左闭右开，注意终止时不要多跑了
// d)常用双指针来做，每个循环都要检查两个指针是否需要变化
// d)注意其中一条循环完毕之后后一条的剩余节点要循环完

// 链表节点连接调整基本套路：
// ReverseBetween 范围反转

// ListNode 单链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseRange 原地反转(from, to)之间的链表，不包含from/to
func ReverseRange(from *ListNode, to *ListNode) {
	cur := from.Next
	for cur != nil && cur.Next != to {
		next := cur.Next         // 2
		cur.Next = cur.Next.Next // 13
		next.Next = from.Next    // 213
		from.Next = next         // h213
	}
}

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

// SortList 链表排序，采用非原地快排，注意点：
// 1.基准选择头部节点，且要单独出来
// 2.新分组要切换队尾与原链表的关联
// 3.合并的时候要记得切断基准节点与原链表的关联
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	sl := &ListNode{}
	ll := &ListNode{}
	sn := sl // 小端当前节点
	ln := ll // 大端当前节点
	// 单次比较，比较基准取头节点
	cur := head.Next // 注意基准要单独出来，否则不能满足一个元素的退出条件
	for cur != nil {
		if cur.Val < head.Val {
			sn.Next = cur
			sn = cur
		} else {
			ln.Next = cur
			ln = cur
		}
		cur = cur.Next
	}
	// 切断原链表
	sn.Next = nil
	ln.Next = nil
	// 递归快排
	sl = SortList(sl.Next)
	ll = SortList(ll.Next)
	// 合并返回新队列
	cur = sl
	if cur != nil {
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = head
		head.Next = nil // 注意要切断原head
		if ll != nil {
			head.Next = ll
		}
		return sl
	}
	head.Next = ll
	return head
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
