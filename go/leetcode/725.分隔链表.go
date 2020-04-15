/*
 * @lc app=leetcode.cn id=725 lang=golang
 *
 * [725] 分隔链表
 *
 * https://leetcode-cn.com/problems/split-linked-list-in-parts/description/
 *
 * algorithms
 * Medium (48.65%)
 * Likes:    23
 * Dislikes: 0
 * Total Accepted:    1.8K
 * Total Submissions: 3.8K
 * Testcase Example:  '[1,2,3,4]\n5'
 *
 * 给定一个头结点为 root 的链表, 编写一个函数以将链表分隔为 k 个连续的部分。
 * 
 * 每部分的长度应该尽可能的相等: 任意两部分的长度差距不能超过 1，也就是说可能有些部分为 null。
 * 
 * 这k个部分应该按照在链表中出现的顺序进行输出，并且排在前面的部分的长度应该大于或等于后面的长度。
 * 
 * 返回一个符合上述规则的链表的列表。
 * 
 * 举例： 1->2->3->4, k = 5 // 5 结果 [ [1], [2], [3], [4], null ]
 * 
 * 示例 1：
 * 
 * 
 * 输入: 
 * root = [1, 2, 3], k = 5
 * 输出: [[1],[2],[3],[],[]]
 * 解释:
 * 输入输出各部分都应该是链表，而不是数组。
 * 例如, 输入的结点 root 的 val= 1, root.next.val = 2, \root.next.next.val = 3, 且
 * root.next.next.next = null。
 * 第一个输出 output[0] 是 output[0].val = 1, output[0].next = null。
 * 最后一个元素 output[4] 为 null, 它代表了最后一个部分为空链表。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入: 
 * root = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], k = 3
 * 输出: [[1, 2, 3, 4], [5, 6, 7], [8, 9, 10]]
 * 解释:
 * 输入被分成了几个连续的部分，并且每部分的长度相差不超过1.前面部分的长度大于等于后面部分的长度。
 * 
 * 
 * 
 * 
 * 提示:
 * 
 * 
 * root 的长度范围： [0, 1000].
 * 输入的每个节点的大小范围：[0, 999].
 * k 的取值范围： [1, 50].
 * 
 * 
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

func splitListToParts(root *ListNode, k int) []*ListNode {
	ret := make([]*ListNode, k)
	cur := root
	l := 0
	for cur != nil {
		cur = cur.Next
		l++
	}
	per := l / k
	les := l % k
	cur = root
	for i := 0; i < k; i++ {
		ret[i] = cur
		pre := &ListNode{Next: cur}
		j := 0
		for pre.Next != nil && j < per {
			pre = pre.Next
			j++
		}
		if pre.Next != nil && les > 0 {
			pre = pre.Next
			les--
		}
		tmp := pre.Next
		pre.Next = nil
		cur = tmp
	}
	return ret
}

