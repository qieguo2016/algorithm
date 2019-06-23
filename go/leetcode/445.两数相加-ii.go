/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 *
 * https://leetcode-cn.com/problems/add-two-numbers-ii/description/
 *
 * algorithms
 * Medium (48.63%)
 * Likes:    63
 * Dislikes: 0
 * Total Accepted:    4.2K
 * Total Submissions: 8.5K
 * Testcase Example:  '[7,2,4,3]\n[5,6,4]'
 *
 * 给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。
 * 
 * 
 * 
 * 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 * 
 * 进阶:
 * 
 * 如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
 * 
 * 示例:
 * 
 * 
 * 输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出: 7 -> 8 -> 0 -> 7
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	arr1 := []*ListNode{}
	arr2 := []*ListNode{}
	for c := l1; c != nil; c = c.Next {
		arr1 = append(arr1, c)
	}
	for c := l2; c != nil; c = c.Next {
		arr2 = append(arr2, c)
	}
	len1 := len(arr1)
	len2 := len(arr2)
	dummy := &ListNode{}
	i := 0
	n := 0  // 进位数
	for i < len1 || i < len2 {
		s := n
		if i < len1 {
			s += arr1[len1-1-i].Val
		}
		if i < len2 {
			s += arr2[len2-1-i].Val
		}
		n = s / 10
		c := &ListNode{}
		c.Val = s % 10
		c.Next = dummy.Next
		dummy.Next = c
		i++
	}
	if n != 0 {
		c := &ListNode{}
		c.Val = n % 10
		c.Next = dummy.Next
		dummy.Next = c
	}
	return dummy.Next
}

