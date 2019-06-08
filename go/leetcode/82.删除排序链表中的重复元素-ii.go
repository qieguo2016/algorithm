/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (41.32%)
 * Likes:    109
 * Dislikes: 0
 * Total Accepted:    10.8K
 * Total Submissions: 26.2K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
 * 
 * 示例 1:
 * 
 * 输入: 1->2->3->3->4->4->5
 * 输出: 1->2->5
 * 
 * 
 * 示例 2:
 * 
 * 输入: 1->1->1->2->3
 * 输出: 2->3
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	cur := dummy.Next
	n := 0
	for cur != nil {
		if cur.Next != nil && cur.Next.Val == pre.Next.Val {  
			cur = cur.Next
			n++
			continue
		}
		if n > 0 {
			pre.Next = cur.Next
			cur = cur.Next
			n = 0
			continue
		}
		pre = pre.Next
		cur = cur.Next
	}
	return dummy.Next
}

