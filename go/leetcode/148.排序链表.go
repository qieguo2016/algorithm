/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 *
 * https://leetcode-cn.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (60.39%)
 * Likes:    194
 * Dislikes: 0
 * Total Accepted:    13.4K
 * Total Submissions: 22.3K
 * Testcase Example:  '[4,2,1,3]'
 *
 * 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
 * 
 * 示例 1:
 * 
 * 输入: 4->2->1->3
 * 输出: 1->2->3->4
 * 
 * 
 * 示例 2:
 * 
 * 输入: -1->5->3->4->0
 * 输出: -1->0->3->4->5
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 采用非原地快排
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	sl := &ListNode{}
	ll := &ListNode{}
	sn := sl // 小端当前节点
	ln := ll // 大端当前节点
	// 单次比较，比较基准取头节点
	cur := head.Next // 注意基准要单独出来，否则不能满足一个元素的退出条件
	for cur != nil {
		if cur.Val < head.Val {
			sn.Next = cur
			sn = cur
		} else {
			ln.Next = cur
			ln = cur
		}
		cur = cur.Next
	}
	// 切断原链表
	sn.Next = nil
	ln.Next = nil
	// 递归快排
	sl = sortList(sl.Next)
	ll = sortList(ll.Next)
	// 合并返回新队列
	cur = sl
	if cur != nil {
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = head
		head.Next = nil  // 注意要切断原head
		if ll != nil {
			head.Next = ll
		}
		return sl
	}
	head.Next = ll
	return head
}
