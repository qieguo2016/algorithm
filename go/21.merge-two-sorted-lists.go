/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (51.82%)
 * Total Accepted:    40.7K
 * Total Submissions: 78.3K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 * 示例：
 *
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	cn := &ListNode{}
	ret := &ListNode{Next: cn} // 预加一个链表头，方便遍历
	for {
		// 其中一条链表已经遍历完，直接将另外一条非空链接上去
		if l1 == nil || l2 == nil {
			if l1 != nil {
				cn.Val = l1.Val
				cn.Next = l1.Next
			}
			if l2 != nil {
				cn.Val = l2.Val
				cn.Next = l2.Next
			}
			break
		}

		var v int
		if l1.Val < l2.Val {
			v = l1.Val
			l1 = l1.Next
		} else {
			v = l2.Val
			l2 = l2.Next
		}

		cn.Val = v
		cn.Next = &ListNode{}
		cn = cn.Next
	}
	return ret.Next // 去掉预加的链表头
}
