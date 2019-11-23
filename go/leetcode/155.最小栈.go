/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (49.91%)
 * Likes:    285
 * Dislikes: 0
 * Total Accepted:    45.8K
 * Total Submissions: 91.3K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
 * 
 * 
 * push(x) -- 将元素 x 推入栈中。
 * pop() -- 删除栈顶的元素。
 * top() -- 获取栈顶元素。
 * getMin() -- 检索栈中的最小元素。
 * 
 * 
 * 示例:
 * 
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 * 
 * 
 */

// 最小值也是一个栈，每次发现有更小的值就加入栈，出栈的时候检查是否最小值，是的话从min栈里出去
type MinStack struct {
	Data []int
	min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
		Data: []int{},
		min: []int{},
	}
}


func (this *MinStack) Push(x int)  {
	this.Data = append(this.Data, x)  // 2 0 3
	if len(this.min) <= 0 || this.min[len(this.min)-1] >= x {  // NOTE: 注意=也要加进去
		this.min = append(this.min, x) // 2 0
	}
}


func (this *MinStack) Pop()  {
	if len(this.Data) <= 0 {
		return
	}
	v := this.Data[len(this.Data)-1]  // 3 0 2
	this.Data = this.Data[:len(this.Data)-1]
	if len(this.min) > 0 && v <= this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1] // 2 0 
	}
}


func (this *MinStack) Top() int {
	if len(this.Data) <= 0 {
		return 0
	}
	return this.Data[len(this.Data)-1]
}


func (this *MinStack) GetMin() int {
    if len(this.min) <= 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

