/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode-cn.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (51.73%)
 * Likes:    205
 * Dislikes: 0
 * Total Accepted:    11.7K
 * Total Submissions: 22.5K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
 * 
 * k 是一个正整数，它的值小于或等于链表的长度。
 * 
 * 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 * 
 * 示例 :
 * 
 * 给定这个链表：1->2->3->4->5
 * 
 * 当 k = 2 时，应当返回: 2->1->4->3->5
 * 
 * 当 k = 3 时，应当返回: 3->2->1->4->5
 * 
 * 说明 :
 * 
 * 
 * 你的算法只能使用常数的额外空间。
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 * 
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 原地反转(from, to)之间的链表，不包含from/to
func reverseRange(from *ListNode, to *ListNode) {
	cur := from.Next
	for cur != nil && cur.Next != to {
		next := cur.Next  // 2
		cur.Next = cur.Next.Next  // 13
		next.Next = from.Next  // 213
		from.Next = next  // h213
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	curHead := dummy  // h
	cur := dummy.Next  // 1
	n := 1
	for cur != nil {  // 不判断 cur.Next != nil，为了多走一步
		cur = cur.Next  // 2, nil
		if n % k == 0 {  // 1
			tmp := curHead.Next
			reverseRange(curHead, cur)  // 
			curHead = tmp  // 
		}
		n++  // 2
	}
	return dummy.Next
}

