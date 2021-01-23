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

// func reverseList(head *ListNode) *ListNode {
// 	dummy := &ListNode{}
// 	dummy.Next = head
// 	for head != nil && head.Next != nil {
// 		n := head.Next
// 		head.Next = head.Next.Next
// 		n.Next = dummy.Next
// 		dummy.Next = n
// 	}
// 	return dummy.Next
// }

// 递归实现，借助函数调用栈实现
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList(head.Next) // 用递归找到最后一个
	head.Next.Next = head // 4>3  3>2
	head.Next = nil // 3>nil 2>nil
	return last
}
