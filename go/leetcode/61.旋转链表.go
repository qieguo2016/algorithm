/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 *
 * https://leetcode-cn.com/problems/rotate-list/description/
 *
 * algorithms
 * Medium (38.17%)
 * Likes:    101
 * Dislikes: 0
 * Total Accepted:    15.5K
 * Total Submissions: 40.6K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
 * 
 * 示例 1:
 * 
 * 输入: 1->2->3->4->5->NULL, k = 2
 * 输出: 4->5->1->2->3->NULL
 * 解释:
 * 向右旋转 1 步: 5->1->2->3->4->NULL
 * 向右旋转 2 步: 4->5->1->2->3->NULL
 * 
 * 
 * 示例 2:
 * 
 * 输入: 0->1->2->NULL, k = 4
 * 输出: 2->0->1->NULL
 * 解释:
 * 向右旋转 1 步: 2->0->1->NULL
 * 向右旋转 2 步: 1->2->0->NULL
 * 向右旋转 3 步: 0->1->2->NULL
 * 向右旋转 4 步: 2->0->1->NULL
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 两种解法
// 最优解：一次遍历将链表存入数组中，然后做转换，空间换时间
// 其他：类似删除倒数第k个类似，当发现k大于链表长度则除余之后递归调用
func rotateRightV1(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k <= 0 {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	ch := dummy
	c := ch.Next
	n := 1
	for c.Next != nil && n < k  {
		c = c.Next
		n++
	}
	// c=2, n=2
	for c.Next != nil {
		ch = ch.Next
		c = c.Next
		n++
	}
	// c=5, h=3, n=5
	// s := k % n  // 2 % 5 = 2
	if k < n {
		dummy.Next = ch.Next  // h>45
		ch.Next = nil   // 3>nil
		c.Next = head  // h45>123>nil
	} else {
		return rotateRight(head, k % n)
	}
	return dummy.Next
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	if k<=0 {
		return head
	}
	// 一遍遍历将链表用数组存起来，然后用k%n得到要移动的点进行移动
	arr := []*ListNode{}
	for c := head; c != nil; c = c.Next {
		arr = append(arr, c)
	}

	l := len(arr)  // 2
	s := k % l   // 1
	if s == 0 {
		return head
	}
	dummy := &ListNode{}
	start := arr[l-1-s]  // 1
	end := arr[l-1]  // 5
	dummy.Next = start.Next
	start.Next = nil   // 3>nil
	end.Next = arr[0]  // h45>123>nil
	return dummy.Next
}

