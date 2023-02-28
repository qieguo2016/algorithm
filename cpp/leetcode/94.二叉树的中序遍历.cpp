/*
 * @lc app=leetcode.cn id=94 lang=cpp
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode.cn/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Easy (76.23%)
 * Likes:    1702
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,2,3]
 * 输出：[1,3,2]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [0, 100] 内
 * -100 <= Node.val <= 100
 *
 *
 *
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 *
 */

#include <stack>
#include <vector>

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
class Solution {
private:
  void inorderTraversalRecurrent(TreeNode *root, std::vector<int> &res) {
    if (root == nullptr) {
      return;
    }
    inorderTraversalRecurrent(root->left, res);
    res.push_back(root->val);
    inorderTraversalRecurrent(root->right, res);
  }

  void inorderTraversalIterative(TreeNode *root, std::vector<int> &res) {
    if (root == nullptr) {
      return;
    }

    std::stack<TreeNode *> stack;
    TreeNode *cur = root;

    while (cur != nullptr || !stack.empty()) {
      if (cur != nullptr) {
        stack.push(cur);
        cur = cur->left;
      } else {

        cur = stack.top();
        stack.pop();
        res.push_back(cur->val);
        cur = cur->right;
      }
    }
  }

public:
  std::vector<int> inorderTraversal(TreeNode *root) {
    std::vector<int> res;
    // inorderTraversalRecurrent(root, res);
    inorderTraversalIterative(root, res);
    return res;
  }
};
// @lc code=end
