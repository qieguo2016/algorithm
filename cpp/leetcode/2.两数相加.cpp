/*
 * @lc app=leetcode.cn id=2 lang=cpp
 *
 * [2] 两数相加
 *
 * https://leetcode-cn.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (40.28%)
 * Likes:    6309
 * Dislikes: 0
 * Total Accepted:    865.8K
 * Total Submissions: 2.1M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给你两个 非空
 * 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
 *
 * 请你将两个数相加，并以相同形式返回一个表示和的链表。
 *
 * 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：l1 = [2,4,3], l2 = [5,6,4]
 * 输出：[7,0,8]
 * 解释：342 + 465 = 807.
 *
 *
 * 示例 2：
 *
 *
 * 输入：l1 = [0], l2 = [0]
 * 输出：[0]
 *
 *
 * 示例 3：
 *
 *
 * 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
 * 输出：[8,9,9,9,0,0,0,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 每个链表中的节点数在范围 [1, 100] 内
 * 0
 * 题目数据保证列表表示的数字不含前导零
 *
 *
 */

#include <iostream>
#include <vector>
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
  ListNode *addTwoNumbers(ListNode *l1, ListNode *l2) {
    ListNode *dummy = new ListNode();
    ListNode *a = l1;
    ListNode *b = l2;
    ListNode *c = dummy;
    int n = 0;
    while (a != nullptr || b != nullptr || n > 0) {
      ListNode *cur = new ListNode(n);
      if (a != nullptr) {
        cur->val += a->val;
        a = a->next;
      }
      if (b != nullptr) {
        cur->val += b->val;
        b = b->next;
      }
      if (cur->val >= 10) {
        cur->val -= 10;
        n = 1;
      } else {
        n = 0;
      }
      c->next = cur;
      c = c->next;
    }
    return dummy->next;
  }
};
// @lc code=end

ListNode *make_list(std::vector<int> input) {
  ListNode *head = new ListNode();
  ListNode *cur = head;
  for (int i = 0; i < input.size(); i++) {
    cur->next = new ListNode(input[i]);
    cur = cur->next;
  }
  return head;
}

void print_ret(ListNode *ret) {
  ListNode *cur = ret;
  std::cout << "ret: ";
  while (cur != nullptr) {
    std::cout << cur->val << ", ";
    cur = cur->next;
  }
  std::cout << std::endl;
}

int main(int argc, const char **argv) {
  Solution s;
  ListNode *a = make_list({9, 9, 9, 9, 9, 9, 9});
  ListNode *b = make_list({9, 9, 9, 9});
  auto ret = s.addTwoNumbers(a, b);
  print_ret(ret);
  return 0;
}
