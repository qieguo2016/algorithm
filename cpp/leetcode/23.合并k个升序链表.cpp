/*
 * @lc app=leetcode.cn id=23 lang=cpp
 *
 * [23] 合并K个升序链表
 *
 * https://leetcode-cn.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (55.28%)
 * Likes:    1342
 * Dislikes: 0
 * Total Accepted:    268.7K
 * Total Submissions: 486K
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 给你一个链表数组，每个链表都已经按升序排列。
 *
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 *
 *
 *
 * 示例 1：
 *
 * 输入：lists = [[1,4,5],[1,3,4],[2,6]]
 * 输出：[1,1,2,3,4,4,5,6]
 * 解释：链表数组如下：
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * 将它们合并到一个有序链表中得到。
 * 1->1->2->3->4->4->5->6
 *
 *
 * 示例 2：
 *
 * 输入：lists = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 * 输入：lists = [[]]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] 按 升序 排列
 * lists[i].length 的总和不超过 10^4
 *
 *
 */

// Definition for singly-linked list.
#include <queue>
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
private:
  struct Element {
    ListNode *ptr;
    bool operator<(const Element &tar) const {
      return ptr->val > tar.ptr->val; // 小顶堆，升序排序用
    }
  };

public:
  ListNode *mergeKLists(std::vector<ListNode *> &lists) {
    ListNode dummy;
    ListNode *cur = &dummy;
    std::priority_queue<Element> queue;
    for (ListNode *node : lists) {
      if (node != nullptr) {
        queue.push({node});
      }
    }
    while (!queue.empty()) {
      auto el = queue.top();
      queue.pop();
      cur->next = el.ptr;
      cur = cur->next;
      if (el.ptr->next != nullptr) {
        queue.push({el.ptr->next});
      }
    }
    return dummy.next;
  }
};
// @lc code=end
