/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (54.62%)
 * Likes:    222
 * Dislikes: 0
 * Total Accepted:    59.3K
 * Total Submissions: 108.5K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 * 
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 返回锯齿形层次遍历如下：
 * 
 * [
 * ⁠ [3],
 * ⁠ [20,9],
 * ⁠ [15,7]
 * ]
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

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	k := 0
	isReverse := false
	for len(q) > 0 {
		curRet := []int{}
		curQueue := []*TreeNode{}
		for _, node := range q {
			// dequeue
			curRet = append(curRet, node.Val)
			// enqueue
			if node.Left != nil {
				curQueue = append(curQueue, node.Left)
			}
			if node.Right != nil {
				curQueue = append(curQueue, node.Right)
			}
		}
		if isReverse {
			reverse(curRet)
		}
		if k >= len(ret) {
			ret = append(ret, []int{})
		}
		ret[k] = curRet


		isReverse = !isReverse
		q = curQueue
		k++  // level down
	}
	return ret
}
// @lc code=end

