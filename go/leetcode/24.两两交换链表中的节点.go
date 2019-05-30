/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 保存两个全局变量：局部转换头节点和当前节点
func swapPairs(head *ListNode) *ListNode {
		dummy := &ListNode{}
		dummy.Next = head
		curHead := dummy
		cur := dummy.Next
		for cur != nil && cur.Next != nil {
			next := cur.Next  // 2, 4
			cur.Next = next.Next  // 13, 35
			next.Next = cur  // 213, 435
			curHead.Next = next  // h213, 21h435
			curHead = cur  // 21h3, 2143h5
			cur = cur.Next  // 3, 5
		}
		return dummy.Next
}

