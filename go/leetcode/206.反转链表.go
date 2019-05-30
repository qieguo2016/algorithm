/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	for head != nil && head.Next != nil {
		n := head.Next
		head.Next = head.Next.Next
		n.Next = dummy.Next
		dummy.Next = n
	}
	return dummy.Next
}
