/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
 *
 * https://leetcode-cn.com/problems/binary-search-tree-iterator/description/
 *
 * algorithms
 * Medium (75.65%)
 * Likes:    316
 * Dislikes: 0
 * Total Accepted:    36.3K
 * Total Submissions: 48K
 * Testcase Example:  '["BSTIterator","next","next","hasNext","next","hasNext","next","hasNext","next","hasNext"]\n[[[7,3,15,null,null,9,20]],[],[],[],[],[],[],[],[],[]]'
 *
 * 实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。
 * 
 * 调用 next() 将返回二叉搜索树中的下一个最小的数。
 * 
 * 
 * 
 * 示例：
 * 
 * 
 * 
 * BSTIterator iterator = new BSTIterator(root);
 * iterator.next();    // 返回 3
 * iterator.next();    // 返回 7
 * iterator.hasNext(); // 返回 true
 * iterator.next();    // 返回 9
 * iterator.hasNext(); // 返回 true
 * iterator.next();    // 返回 15
 * iterator.hasNext(); // 返回 true
 * iterator.next();    // 返回 20
 * iterator.hasNext(); // 返回 false
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
 * 你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。
 * 
 * 
 */

// @lc code=start

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

type BSTIterator struct {
	root *TreeNode
	stack []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		root: root,
		stack: getPathToMin(root),
	}
}

func getPathToMin(root *TreeNode) []*TreeNode {
	stack := make([]*TreeNode, 0)
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	return stack
}

func (this *BSTIterator) Next() int {
	if len(this.stack) <= 0 {
		return -1
	}
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[0:len(this.stack)-1]
	if node == nil {
		return -1
	}
	if node.Right != nil {
		this.stack = append(this.stack, getPathToMin(node.Right)...)
	}
	return node.Val
}


func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

