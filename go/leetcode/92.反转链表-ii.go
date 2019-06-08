/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 *
 * https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (44.01%)
 * Likes:    137
 * Dislikes: 0
 * Total Accepted:    11K
 * Total Submissions: 24.9K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
 * 
 * 说明:
 * 1 ≤ m ≤ n ≤ 链表长度。
 * 
 * 示例:
 * 
 * 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 * 输出: 1->4->3->2->5->NULL
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// h>1>2>3>4>5
// h>1>3>2>4>5
// h>1>4>3>2>5
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	i := 1
	pre := dummy
	cur := dummy.Next
	for cur != nil && cur.Next != nil {
		if i >= m && i < n {   // 向后调整，不需要i==n
			next := cur.Next
			cur.Next = next.Next
			next.Next = pre.Next
			pre.Next = next
			i++
			continue
		}
		pre = pre.Next
		cur = cur.Next
		i++
	}
	return dummy.Next
}
