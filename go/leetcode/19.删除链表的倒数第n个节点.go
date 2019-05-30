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
	if head == nil {
		return nil
	}
	if n < 0 {
		return head
	}
	k := 1  // n从1开始，匹配
	c := head
	for c.Next != nil && k < n {
		c = c.Next
		k++
	}
	pre := &ListNode{Next: head}  // 从1开始计算，所以要预留一个空位
	for c.Next != nil {
		pre = pre.Next
		c = c.Next
	}
	if n == k {
		if pre.Next == head { // 支持删除头部
			return pre.Next.Next
		}
		node := pre.Next
		pre.Next = pre.Next.Next
		node.Next = nil
		return head
	}
	return nil
}

