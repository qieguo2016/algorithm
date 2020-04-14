/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// n从1开始，仅有一个元素移除返回nil
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 0 {
		return head
	}
	dummy := &ListNode{Next: head}
	k := 0
	left := dummy
	right := dummy
	for right.Next != nil && k < n {
		right = right.Next
		k++
	}
	if k < n {   // n超过链表长度
		return dummy.Next
	}
	for right.Next != nil {
		right = right.Next
		left = left.Next
	}
	node := left.Next
	if node != nil {
		left.Next = node.Next
		node.Next = nil
	}
	return dummy.Next
}

