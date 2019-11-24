/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] 常数时间插入、删除和获取随机元素
 *
 * https://leetcode-cn.com/problems/insert-delete-getrandom-o1/description/
 *
 * algorithms
 * Medium (46.26%)
 * Likes:    58
 * Dislikes: 0
 * Total Accepted:    6.4K
 * Total Submissions: 13.7K
 * Testcase Example:  '["RandomizedSet","insert","remove","insert","getRandom","remove","insert","getRandom"]\n[[],[1],[2],[2],[],[1],[2],[]]'
 *
 * 设计一个支持在平均 时间复杂度 O(1) 下，执行以下操作的数据结构。
 * 
 * 
 * insert(val)：当元素 val 不存在时，向集合中插入该项。
 * remove(val)：元素 val 存在时，从集合中移除该项。
 * getRandom：随机返回现有集合中的一项。每个元素应该有相同的概率被返回。
 * 
 * 
 * 示例 :
 * 
 * 
 * // 初始化一个空的集合。
 * RandomizedSet randomSet = new RandomizedSet();
 * 
 * // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
 * randomSet.insert(1);
 * 
 * // 返回 false ，表示集合中不存在 2 。
 * randomSet.remove(2);
 * 
 * // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
 * randomSet.insert(2);
 * 
 * // getRandom 应随机返回 1 或 2 。
 * randomSet.getRandom();
 * 
 * // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
 * randomSet.remove(1);
 * 
 * // 2 已在集合中，所以返回 false 。
 * randomSet.insert(2);
 * 
 * // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
 * randomSet.getRandom();
 * 
 * 
 */

/*
 * 1. 为了达到O(1)的插入删除，需要使用hash
 * 2. 为了达到O(1)的随机数，需要用数组保存，然后随机取数组长度内的下标
 * 3. 增加数组之后，增加可以直接增加到数组末尾，删除的话可以交换被删元素与数组尾部，这样只需要移动一个元素
 */

// @lc code=start
import (
	"math/rand"
)

type RandomizedSet struct {
	Data map[int]int  // k为传入val，v为下标
	Index []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    return RandomizedSet{
		Data: map[int]int{},
		Index: []int{},
	}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Data[val]; ok {
		return false
	}
	this.Index = append(this.Index, val)
	this.Data[val] = len(this.Index) - 1
	return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.Data[val]
    if !ok {
		return false
	}
	last := len(this.Index) - 1
	if idx < last {  // 与队尾交互
		this.Data[this.Index[last]] = idx
		this.Index[idx] = this.Index[last]
	}
	delete(this.Data, val)
	this.Index = this.Index[:last]
	return true
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.Index))
	return this.Index[idx]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

