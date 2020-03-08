/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (64.09%)
 * Likes:    115
 * Dislikes: 0
 * Total Accepted:    14.2K
 * Total Submissions: 22K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 根据一棵树的中序遍历与后序遍历构造二叉树。
 * 
 * 注意:
 * 你可以假设树中没有重复的元素。
 * 
 * 例如，给出
 * 
 * 中序遍历 inorder = [9,3,15,20,7]
 * 后序遍历 postorder = [9,15,7,20,3]
 * 
 * 返回如下的二叉树：
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法的前提是二叉树中无重复元素，所以相同的值可以认为是同一个元素
// 一般树的题目都优先考虑递归做法，构造二叉树的递归就是找到父节点，然后递归左右子树
// 然后看题目，中序遍历的特点是左父右，后序遍历的特点是左右父
// 所以可以先从后序找到父节点在最后一个节点，然后从中序找到左子树，剩下的就是右子树
func buildTree(inorder []int, postorder []int) *TreeNode {
    return buildTreeHelper(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func buildTreeHelper(inorder []int, iLeft int, iRight int, postorder []int, pLeft int, pRight int) *TreeNode {
    if iLeft > iRight || pLeft > pRight {
		return nil
	}

	root := &TreeNode{Val: postorder[pRight]}
	// 从中序找到左右子树分界
	var i int
	for i = iLeft; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	// 后序是左右根，根据左子树的个数(i-iLeft)可以截取左右子树
	root.Left = buildTreeHelper(inorder, iLeft, i - 1, postorder, pLeft, pLeft + i - iLeft - 1)
	root.Right = buildTreeHelper(inorder, i + 1, iRight, postorder, pLeft + i - iLeft, pRight - 1)
	return root
}
// @lc code=end

