/*
 * @lc app=leetcode.cn id=25 lang=cpp
 *
 * [25] K 个一组翻转链表
 *
 * https://leetcode.cn/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (67.74%)
 * Likes:    1907
 * Dislikes: 0
 * Total Accepted:    430.9K
 * Total Submissions: 636.1K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
 *
 * k
 * 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], k = 2
 * 输出：[2,1,4,3,5]
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：head = [1,2,3,4,5], k = 3
 * 输出：[3,2,1,4,5]
 *
 *
 *
 * 提示：
 *
 *
 * 链表中的节点数目为 n
 * 1 <= k <= n <= 5000
 * 0 <= Node.val <= 1000
 *
 *
 *
 *
 * 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
 *
 *
 *
 *
 */

// Definition for singly-linked list.
#include <sys/_types/_wint_t.h>
struct ListNode {
  int val;
  ListNode *next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode *next) : val(x), next(next) {}
};

// @lc code=start
class Solution {
public:
  ListNode *reverseRange(ListNode *start, ListNode *end) {
    ListNode *cur = start->next, *next;
    while (cur->next != nullptr && cur->next != end) {
      next = cur->next; // s sn --- c n nn --- e  => s n sn --- c nn --- e
      cur->next = next->next;
      next->next = start->next;
      start->next = next;
    }
    return cur; // 下一个start
  }
  ListNode *reverseKGroup(ListNode *head, int k) {
    ListNode dummy;
    dummy.next = head;
    int i = 0;
    ListNode *start = &dummy, *end = dummy.next;
    while (end != nullptr) {
      end = end->next;
      i++;
      if (i >= k) {
        start = reverseRange(start, end);
        i = 0;
      }
    }

    return dummy.next;
  }
};
// @lc code=end
