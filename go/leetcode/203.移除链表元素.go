/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 *
 * https://leetcode-cn.com/problems/remove-linked-list-elements/description/
 *
 * algorithms
 * Easy (41.18%)
 * Likes:    238
 * Dislikes: 0
 * Total Accepted:    25.9K
 * Total Submissions: 62.8K
 * Testcase Example:  '[1,2,6,3,4,5,6]\n6'
 *
 * 删除链表中等于给定值 val 的所有节点。
 * 
 * 示例:
 * 
 * 输入: 1->2->6->3->4->5->6, val = 6
 * 输出: 1->2->3->4->5
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
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	for pre != nil && pre.Next != nil {
		if pre.Next.Val == val {
			tmp := pre.Next
			pre.Next = tmp.Next
			tmp.Next = nil
		} else {
			pre = pre.Next
		}
	}
	return dummy.Next
}

