/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 *
 * https://leetcode-cn.com/problems/odd-even-linked-list/description/
 *
 * algorithms
 * Medium (56.54%)
 * Likes:    62
 * Dislikes: 0
 * Total Accepted:    10.5K
 * Total Submissions: 18.6K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
 * 
 * 请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。
 * 
 * 示例 1:
 * 
 * 输入: 1->2->3->4->5->NULL
 * 输出: 1->3->5->2->4->NULL
 * 
 * 
 * 示例 2:
 * 
 * 输入: 2->1->3->5->6->4->7->NULL 
 * 输出: 2->3->6->7->1->5->4->NULL
 * 
 * 说明:
 * 
 * 
 * 应当保持奇数节点和偶数节点的相对顺序。
 * 链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
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
func oddEvenList(head *ListNode) *ListNode {
	// h123456
	// h132456
	// h135246
	// h13572468
	// 两个指针分别指向奇偶队尾
	// i表示序列号，用以判断奇偶
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy.Next
	oddEnd := dummy.Next
	for cur != nil && cur.Next != nil && cur.Next.Next != nil {
		c1 := cur.Next  // 2, 4
		c2 := cur.Next.Next   // 3, 5
		evenStart := oddEnd.Next // 2, 2
		oddEnd.Next = c2  // 1>3, 135
		c1.Next = c2.Next  // 2>4, 4>6
		oddEnd = oddEnd.Next  // 3, 5
		oddEnd.Next = evenStart // 135246
		cur = c1  // 4
	}
	return dummy.Next
}

