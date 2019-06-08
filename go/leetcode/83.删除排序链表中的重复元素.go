/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (45.95%)
 * Likes:    158
 * Dislikes: 0
 * Total Accepted:    28.6K
 * Total Submissions: 62.2K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
 * 
 * 示例 1:
 * 
 * 输入: 1->1->2
 * 输出: 1->2
 * 
 * 
 * 示例 2:
 * 
 * 输入: 1->1->2->3->3
 * 输出: 1->2->3
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
				cur = cur.Next   // 1>1 
				n++		// 1
				continue
			}
			if n > 0 {   // 1 
				pre.Next = cur  // p>1>2
				pre = pre.Next			// 1
				cur = cur.Next   // 2
				n = 0
				continue
			}
			pre = pre.Next
			cur = cur.Next
			continue
		}
		return dummy.Next
}

