/*
 * @lc app=leetcode.cn id=109 lang=cpp
 *
 * [109] 有序链表转换二叉搜索树
 *
 * https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/description/
 *
 * algorithms
 * Medium (76.46%)
 * Likes:    826
 * Dislikes: 0
 * Total Accepted:    145K
 * Total Submissions: 189.6K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 给定一个单链表的头节点  head ，其中的元素 按升序排序
 * ，将其转换为高度平衡的二叉搜索树。
 *
 * 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差不超过
 * 1。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 *
 * 输入: head = [-10,-3,0,5,9]
 * 输出: [0,-3,9,-10,null,5]
 * 解释:
 * 一个可能的答案是[0，-3,9，-10,null,5]，它表示所示的高度平衡的二叉搜索树。
 *
 *
 * 示例 2:
 *
 *
 * 输入: head = []
 * 输出: []
 *
 *
 *
 *
 * 提示:
 *
 *
 * head 中的节点数在[0, 2 * 10^4] 范围内
 * -10^5 <= Node.val <= 10^5
 *
 *
 */

// Definition for singly-linked list.
struct ListNode {
  int val;
  ListNode *next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode *next) : val(x), next(next) {}
};

// Definition for a binary tree node.
struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
  TreeNode(int x, TreeNode *left, TreeNode *right)
      : val(x), left(left), right(right) {}
};

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left),
 * right(right) {}
 * };
 */
class Solution {
public:
  // 二叉树中序遍历刚好就是链表，也就是从中序遍历结果反推二叉树
  // 可以先构建root节点，待到中序遍历到该位置再填入
  TreeNode *sortedListToBST(ListNode *head) {
    return build(head, 0, len(head)-1);
  }

private:
  int len(ListNode *head) {
    int ret = 0;
    while (head != nullptr) {
      head = head->next;
      ret++;
    }
    return ret;
  }
  // head 是指针引用，为了在内部更新head
  TreeNode *build(ListNode *&head, int l, int r) {
    if (l > r) {
      return nullptr;
    }
    int mid = (l + r + 1) / 2;
    TreeNode *root = new TreeNode();
    root->left = build(head, l, mid - 1);
    root->val = head->val;
    head = head->next;
    root->right = build(head, mid + 1, r);
    return root;
  }
};
// @lc code=end
