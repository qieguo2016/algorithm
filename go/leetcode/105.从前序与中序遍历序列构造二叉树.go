/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (61.23%)
 * Likes:    247
 * Dislikes: 0
 * Total Accepted:    24.7K
 * Total Submissions: 40.1K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 * 
 * 注意:
 * 你可以假设树中没有重复的元素。
 * 
 * 例如，给出
 * 
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inorder = [9,3,15,20,7]
 * 
 * 返回如下的二叉树：
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 前序是根左右的遍历方式，中序是左根右的遍历方式
// 3是根节点，然后以根结点将中序一分为2得到左右子树，按照左右子树个数将前序也分割成两个子树
// 分别递归左右子树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 || len(inorder) <= 0 {
		return nil
	}
	mid := 0
	for mid < len(inorder) && inorder[mid] != preorder[0] {
		mid++
	}
	ret := &TreeNode{Val: preorder[0]}
	ret.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	ret.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return ret
}

