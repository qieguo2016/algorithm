// package leetcode

/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 *
 * https://leetcode-cn.com/problems/implement-stack-using-queues/description/
 *
 * algorithms
 * Easy (61.24%)
 * Likes:    89
 * Dislikes: 0
 * Total Accepted:    22.2K
 * Total Submissions: 36.3K
 * Testcase Example:  '["MyStack","push","push","top","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * 使用队列实现栈的下列操作：
 *
 *
 * push(x) -- 元素 x 入栈
 * pop() -- 移除栈顶元素
 * top() -- 获取栈顶元素
 * empty() -- 返回栈是否为空
 *
 *
 * 注意:
 *
 *
 * 你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty
 * 这些操作是合法的。
 * 你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
 * 你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。
 *
 *
 */

// @lc code=start
import (
	"container/list"
)

type MyStack struct {
	Data *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		Data: list.New(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Data.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	el := this.Data.Back()
	if el == nil {
		return 0
	}
	return this.Data.Remove(el).(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	el := this.Data.Back()
	if el == nil {
		return 0
	}
	return el.Value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Data.Len() <= 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
