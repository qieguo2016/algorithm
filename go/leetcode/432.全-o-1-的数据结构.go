/*
 * @lc app=leetcode.cn id=432 lang=golang
 *
 * [432] 全 O(1) 的数据结构
 *
 * https://leetcode-cn.com/problems/all-oone-data-structure/description/
 *
 * algorithms
 * Hard (35.21%)
 * Likes:    25
 * Dislikes: 0
 * Total Accepted:    1.8K
 * Total Submissions: 5.1K
 * Testcase Example:  '["AllOne","getMaxKey","getMinKey"]\n[[],[],[]]'
 *
 * 实现一个数据结构支持以下操作：
 *
 *
 * Inc(key) - 插入一个新的值为 1 的 key。或者使一个存在的 key 增加一，保证 key 不为空字符串。
 * Dec(key) - 如果这个 key 的值是 1，那么把他从数据结构中移除掉。否者使一个存在的 key 值减一。如果这个 key
 * 不存在，这个函数不做任何事情。key 保证不为空字符串。
 * GetMaxKey() - 返回 key 中值最大的任意一个。如果没有元素存在，返回一个空字符串""。
 * GetMinKey() - 返回 key 中值最小的任意一个。如果没有元素存在，返回一个空字符串""。
 *
 *
 * 挑战：以 O(1) 的时间复杂度实现所有操作。
 *
 */

/*
 * 这个题目与155最小栈那题有点类似，但是有点不同的是每次加减1，分析一下：
 * 1. 必然需要一个hash结构，key输入的key，value中除数值外，一般可以带上一个指针指向其他数据结构的节点
 * 2. 需要保存最大最小值，注意这不是一个值，而是动态变化的一系列值，某个最大值变小之后要将次大的推上来，
 *    那么可选数组、链表等各种结构，hash的值指向这个节点，帮助快速定位
 * 3. 另外一个点是加减1这个操作，加减1之后元素可能前移可能后移，移动的时候都是跨过相等的一批值，
 *    如果将相同值的元素放在同一个节点上，每个节点是一个层的概念，那么移动的话只需要移动到前后层即可。
 *    这时候引入一个新问题，就是如何快速定位到目标点上，想当然就是hash的value中保存目标点的指针，
 *    另外同一层内没有移动操作，只有O(1)的查找，所以每层只需要一个hash保存即可
 */

// @lc code=start
import (
	"container/list"
)

type Node struct {
	Key   string
	Value int
	Level *Level
}
type Level struct {
	Value   int
	Nodes   map[string]bool
	Element *list.Element // 快速定位到链表节点，用以查找前后节点
}

type AllOne struct {
	Data map[string]*Node
	Rank *list.List
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		Data: map[string]*Node{},
		Rank: list.New(),
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	node, ok := this.Data[key]
	if !ok {
		needInsert := false
		var level *Level
		front := this.Rank.Front()
		// 空或者队头value大于1都需要插入
		if front == nil {
			needInsert = true
		} else if level = front.Value.(*Level); level.Value > 1 {
			needInsert = true
		}
		if needInsert {
			level = &Level{Value: 1, Nodes: map[string]bool{}}
			el := this.Rank.PushFront(level)
			level.Element = el
		}
		level.Nodes[key] = true
		node := &Node{Key: key, Value: 1, Level: level}
		this.Data[key] = node
		return
	}
	node.Value++
	originLevel := node.Level
	needInsert := false
	var nextLevel *Level
	nextLevelElement := originLevel.Element.Next()
	// 队尾或者下一个的value不是目标值，需要插入新节点
	if nextLevelElement == nil {
		needInsert = true
	} else if nextLevel = nextLevelElement.Value.(*Level); nextLevel.Value > node.Value {
		needInsert = true
	}
	if needInsert {
		nextLevel = &Level{Value: node.Value, Nodes: map[string]bool{}}
		el := this.Rank.InsertAfter(nextLevel, originLevel.Element)
		nextLevel.Element = el
	}

	// 设置到新level
	nextLevel.Nodes[node.Key] = true
	node.Level = nextLevel
	// 清理原level
	delete(originLevel.Nodes, node.Key)
	if len(originLevel.Nodes) <= 0 {
		this.Rank.Remove(originLevel.Element)
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	node, ok := this.Data[key]
	if !ok {
		return
	}
	node.Value--
	if node.Value <= 0 { // 删除节点
		delete(this.Data, key)
		delete(node.Level.Nodes, key)
		if len(node.Level.Nodes) <= 0 {
			this.Rank.Remove(node.Level.Element)
		}
		return
	}
	originLevel := node.Level
	prevLevelElement := originLevel.Element.Prev()

	needInsert := false
	var prevLevel *Level
	if prevLevelElement == nil {
		needInsert = true
	} else if prevLevel = prevLevelElement.Value.(*Level); prevLevel.Value < node.Value {
		needInsert = true
	}
	if needInsert {
		prevLevel = &Level{Value: node.Value, Nodes: map[string]bool{}}
		el := this.Rank.InsertBefore(prevLevel, originLevel.Element)
		prevLevel.Element = el
	}
	// 设置到新level
	prevLevel.Nodes[node.Key] = true
	node.Level = prevLevel
	// 清理原level
	delete(originLevel.Nodes, node.Key)
	if len(originLevel.Nodes) <= 0 {
		this.Rank.Remove(originLevel.Element)
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	maxLevelElement := this.Rank.Back()
	if maxLevelElement == nil {
		return ""
	}
	maxLevel := maxLevelElement.Value.(*Level)
	for k := range maxLevel.Nodes {
		return k
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	minLevelElement := this.Rank.Front()
	if minLevelElement == nil {
		return ""
	}
	minLevel := minLevelElement.Value.(*Level)
	for k := range minLevel.Nodes {
		return k
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
// @lc code=end
