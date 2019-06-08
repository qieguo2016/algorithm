/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 *
 * https://leetcode-cn.com/problems/partition-list/description/
 *
 * algorithms
 * Medium (49.13%)
 * Likes:    92
 * Dislikes: 0
 * Total Accepted:    8.7K
 * Total Submissions: 17.7K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
 * 
 * 你应当保留两个分区中每个节点的初始相对位置。
 * 
 * 示例:
 * 
 * 输入: head = 1->4->3->2->5->2, x = 3
 * 输出: 1->2->2->4->3->5
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
// 也是快慢指针：找到>=x的点之后，比x小的点插入到慢指针之后
// 1. has_div表示是否找到>=x的点
// 2. pre指针表示可被插入的位置
// 3. 小于x的点都保持在pre点之后，pre指针总是会移动，要么移动到下一个，要不就是移动到新插入点上
func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	has_div := false
	pre := dummy
	cur := dummy
	for cur != nil && cur.Next != nil {
		next := cur.Next   // 因为需要插入，所以取前指针比较合适
		if next.Val >= x {
			if !has_div {    // 第一次，找到分隔点
				has_div = true
			}
			cur = next   // 继续往下走
			continue
		}
		if has_div {   // 有分隔点，那就插入到分隔点
			cur.Next = next.Next  // 从当前位置断开
			next.Next = pre.Next  // 后接pre后指针
			pre.Next = next   // 插入到pre之后
		}
		pre = pre.Next   // 移动pre
		cur = next    // 移动cur
	}
	return dummy.Next
}

