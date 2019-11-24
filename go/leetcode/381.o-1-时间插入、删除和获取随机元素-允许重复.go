// package leetcode

/*
 * @lc app=leetcode.cn id=381 lang=golang
 *
 * [381] O(1) 时间插入、删除和获取随机元素 - 允许重复
 *
 * https://leetcode-cn.com/problems/insert-delete-getrandom-o1-duplicates-allowed/description/
 *
 * algorithms
 * Hard (34.86%)
 * Likes:    31
 * Dislikes: 0
 * Total Accepted:    1.4K
 * Total Submissions: 3.7K
 * Testcase Example:  '["RandomizedCollection","insert","insert","insert","getRandom","remove","getRandom"]\n[[],[1],[1],[2],[],[1],[]]'
 *
 * 设计一个支持在平均 时间复杂度 O(1) 下， 执行以下操作的数据结构。
 *
 * 注意: 允许出现重复元素。
 *
 *
 * insert(val)：向集合中插入元素 val。
 * remove(val)：当 val 存在时，从集合中移除一个 val。
 * getRandom：从现有集合中随机获取一个元素。每个元素被返回的概率应该与其在集合中的数量呈线性相关。
 *
 *
 * 示例:
 *
 * // 初始化一个空的集合。
 * RandomizedCollection collection = new RandomizedCollection();
 *
 * // 向集合中插入 1 。返回 true 表示集合不包含 1 。
 * collection.insert(1);
 *
 * // 向集合中插入另一个 1 。返回 false 表示集合包含 1 。集合现在包含 [1,1] 。
 * collection.insert(1);
 *
 * // 向集合中插入 2 ，返回 true 。集合现在包含 [1,1,2] 。
 * collection.insert(2);
 *
 * // getRandom 应当有 2/3 的概率返回 1 ，1/3 的概率返回 2 。
 * collection.getRandom();
 *
 * // 从集合中删除 1 ，返回 true 。集合现在包含 [1,2] 。
 * collection.remove(1);
 *
 * // getRandom 应有相同概率返回 1 和 2 。
 * collection.getRandom();
 *
 *
 */

// @lc code=start
import (
	"math/rand"
)

type RandomizedCollection struct {
	Data  map[int]map[int]bool // 用map存位置
	Index []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		Data:  map[int]map[int]bool{},
		Index: []int{},
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	notFound := false
	if _, ok := this.Data[val]; !ok {
		this.Data[val] = map[int]bool{}
		notFound = true
	}
	this.Index = append(this.Index, val)
	this.Data[val][len(this.Index)-1] = true
	return notFound
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	dict, ok := this.Data[val]
	if !ok {
		return false
	}
	var target int
	for k := range dict {
		target = k
		break
	}
	if len(dict) <= 1 { // 清空
		delete(this.Data, val)
	} else {
		delete(dict, target)
	}
	lastIdx := len(this.Index) - 1
	lastVal := this.Index[lastIdx]
	if target < lastIdx {
		// 队尾的值挪到目标位然后去掉队尾
		this.Index[target] = lastVal
	}
	this.Index = this.Index[:lastIdx]
	lastDict := this.Data[lastVal]
	if len(lastDict) > 0 {
		lastDict[target] = true
		delete(lastDict, lastIdx)
	}
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	idx := rand.Intn(len(this.Index))
	return this.Index[idx]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
