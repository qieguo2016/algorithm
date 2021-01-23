/*
 * @lc app=leetcode.cn id=19 lang=java
 *
 * [19] 删除链表的倒数第N个节点
 *
 * https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (33.67%)
 * Likes:    394
 * Dislikes: 0
 * Total Accepted:    45.8K
 * Total Submissions: 135.9K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
 * 
 * 示例：
 * 
 * 给定一个链表: 1->2->3->4->5, 和 n = 2.
 * 
 * 当删除了倒数第二个节点后，链表变为 1->2->3->5.
 * 
 * 
 * 说明：
 * 
 * 给定的 n 保证是有效的。
 * 
 * 进阶：
 * 
 * 你能尝试使用一趟扫描实现吗？
 * 
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    // public class ListNode {
    //     int val;
    //     ListNode next;
    //     ListNode(int x) { val = x; }
    // }
    public ListNode removeNthFromEnd(ListNode head, int n) {
        if (head == null || n <= 0) {
            return head;
        }
        int k = 0;
        ListNode dummy = new ListNode(0);
        dummy.next = head;
        ListNode left = dummy;
        ListNode right = dummy;
        while (k < n && right.next != null) {
            right = right.next;
            k++;
        }
        if (k < n) {
            return dummy.next;
        }
        while (right.next != null) {
            left = left.next;
            right = right.next;
        }
        if (left.next != null) {
            ListNode node = left.next;
            left.next = node.next;
            node.next = null;
        }
        return dummy.next;
    }
}
// @lc code=end

