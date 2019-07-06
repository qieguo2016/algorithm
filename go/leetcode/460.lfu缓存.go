/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU缓存
 *
 * https://leetcode-cn.com/problems/lfu-cache/description/
 *
 * algorithms
 * Hard (30.93%)
 * Likes:    25
 * Dislikes: 0
 * Total Accepted:    605
 * Total Submissions: 2K
 * Testcase Example:  '["LFUCache","put","put","get","put","get","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]'
 *
 * 设计并实现最不经常使用（LFU）缓存的数据结构。它应该支持以下操作：get 和 put。
 * 
 * get(key) - 如果键存在于缓存中，则获取键的值（总是正数），否则返回 -1。
 * put(key, value) -
 * 如果键不存在，请设置或插入值。当缓存达到其容量时，它应该在插入新项目之前，使最不经常使用的项目无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，最近最少使用的键将被去除。
 * 
 * 进阶：
 * 你是否可以在 O(1) 时间复杂度内执行两项操作？
 * 
 * 示例：
 * 
 * 
 * LFUCache cache = new LFUCache( 2 ) // capacity (缓存容量)
 * 
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // 返回 1
 * cache.put(3, 3);    // 去除 key 2
 * cache.get(2);       // 返回 -1 (未找到key 2)
 * cache.get(3);       // 返回 3
 * cache.put(4, 4);    // 去除 key 1
 * cache.get(1);       // 返回 -1 (未找到 key 1)
 * cache.get(3);       // 返回 3
 * cache.get(4);       // 返回 4
 * 
 */

// LFUCache 基于内存实现
// 原理：map结构按照kv存储数据，双向链表保存节点使用频率

// LFUChainNode 链表节点，按freq降序排列
type LFUChainNode struct {
	pre   *LFUChainNode
	next  *LFUChainNode
	key   int
	value int
	freq  int
}

// LFUCache 结构
type LFUCache struct {
	capacity int
	length   int
	store    map[int]*LFUChainNode
	head     *LFUChainNode
	tail     *LFUChainNode
}

// Constructor constructor
func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		length:   0,
		store:    map[int]*LFUChainNode{},
	}
}

func (c *LFUCache) adjustNode(node *LFUChainNode) {
	// 更新频率之后判断是否需要往前移动
	for node.pre != nil && node.freq >= node.pre.freq {
		pre1 := node.pre
		pre0 := pre1.pre
		next := node.next

		// 调整元素
		node.pre = pre0
		if pre0 != nil {
			pre0.next = node
		}
		node.next = pre1
		pre1.pre = node
		pre1.next = next
		if next != nil {
			next.pre = pre1
		}
		// 处理链表头部
		if node.pre == nil { // 已经移动到了链表头
			c.head = node
		}
		// 原来是链表尾
		if next == nil {
			c.tail = pre1
		}
	}
}

// Delete 删除key
func (c *LFUCache) Delete(key int) {
	node, exist := c.store[key]
	if !exist {
		return
	}
	delete(c.store, key)
	if c.length == 1 {
		c.head = nil
		c.tail = nil
		c.length--
		return
	}
	// 头结点
	if node.pre == nil {
		c.head = c.head.next
		c.head.pre = nil
	} else {
		node.pre.next = node.next
	}
	// 尾部
	if node.next == nil {
		c.tail = c.tail.pre
		c.tail.next = nil
	} else {
		node.next.pre = node.pre
	}
	c.length--
}

// Get 获取kv，更新使用次数，调整链表顺序
func (c *LFUCache) Get(key int) int {
	node, exist := c.store[key]
	if !exist {
		return -1
	}
	node.freq++
	c.adjustNode(node)
	return node.value
}

// Put 插入
func (c *LFUCache) Put(key int, value int) {
	if c.capacity == 0 {
		return
	}
	node, exist := c.store[key]
	// 已经存在，刷新存储值，刷新频率
	if exist {
		node.value = value
		node.freq++
		c.adjustNode(node)
		return
	}
	// 不存在，新插入
	// 超过容量的时候清理频率最低的key，即队尾
	if c.length+1 > c.capacity {
		c.Delete(c.tail.key)
	}
	node = &LFUChainNode{key: key, value: value, freq: 1}
	c.store[key] = node
	c.length++
	if c.length == 1 {
		c.head = node
		c.tail = node
		return
	}
	// 追加到尾部
	c.tail.next = node
	node.pre = c.tail
	c.tail = node
	c.adjustNode(node)
}
