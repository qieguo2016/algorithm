/*
 * @lc app=leetcode.cn id=105 lang=cpp
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (71.39%)
 * Likes:    1873
 * Dislikes: 0
 * Total Accepted:    460.6K
 * Total Submissions: 645.1K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历，
 * inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
 * 输出: [3,9,20,null,null,15,7]
 *
 *
 * 示例 2:
 *
 *
 * 输入: preorder = [-1], inorder = [-1]
 * 输出: [-1]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= preorder.length <= 3000
 * inorder.length == preorder.length
 * -3000 <= preorder[i], inorder[i] <= 3000
 * preorder 和 inorder 均 无重复 元素
 * inorder 均出现在 preorder
 * preorder 保证 为二叉树的前序遍历序列
 * inorder 保证 为二叉树的中序遍历序列
 *
 *
 */

#include <unordered_map>
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
  std::unordered_map<int, int> inorder_index;
  TreeNode *buildTree(std::vector<int> &preorder, int pre_left, int pre_right,
                      std::vector<int> &inorder, int in_left, int in_right) {
    if (pre_left > pre_right) {
      return nullptr;
    }
    int mid = inorder_index[preorder[pre_left]] - in_left;  // count
    TreeNode *root = new TreeNode(preorder[pre_left]);
    root->left = buildTree(preorder, pre_left + 1, pre_left + mid, inorder,
                           in_left, in_left + mid - 1);
    root->right = buildTree(preorder, pre_left + mid + 1, pre_right, inorder,
                            in_left + mid + 1, in_right);
    return root;
  }

public:
  TreeNode *buildTree(std::vector<int> &preorder, std::vector<int> &inorder) {
    int n = inorder.size();
    for (int i = 0; i < n; i++) {
      inorder_index[inorder[i]] = i;
    }
    return buildTree(preorder, 0, n - 1, inorder, 0, n - 1);
  }
};
// @lc code=end
