/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 *
 * https://leetcode-cn.com/problems/lru-cache/description/
 *
 * algorithms
 * Hard (38.35%)
 * Total Accepted:    5.5K
 * Total Submissions: 14K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 *
 * 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
 * 写入数据 put(key, value) -
 * 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。
 *
 * 进阶:
 *
 * 你是否可以在 O(1) 时间复杂度内完成这两种操作？
 *
 * 示例:
 *
 * LRUCache cache = new LRUCache(2);  // 2为缓存容量
 *
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // 返回  1
 * cache.put(3, 3);    // 该操作会使得密钥 2 作废
 * cache.get(2);       // 返回 -1 (未找到)
 * cache.put(4, 4);    // 该操作会使得密钥 1 作废
 * cache.get(1);       // 返回 -1 (未找到)
 * cache.get(3);       // 返回  3
 * cache.get(4);       // 返回  4
 *
 *
 */

type LRUChainNode struct {
	pre   *LRUChainNode
	next  *LRUChainNode
	key   int
	value int
	ts    int32
}

type LRUCache struct {
	capacity int
	length   int
	store    map[int]*LRUChainNode
	head     *LRUChainNode
	tail     *LRUChainNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		length:   0,
		store:    map[int]*LRUChainNode{},
	}
}

func (this *LRUCache) Delete(key int) {
	node, exist := this.store[key]
	if !exist {
		return
	}
	delete(this.store, key)
	if this.length == 1 {
		this.head = nil
		this.tail = nil
		this.length--
		return
	}
	if node.pre == nil {
		this.head = this.head.next
		this.head.pre = nil
	} else {
		node.pre.next = node.next
	}
	if node.next == nil {
		this.tail = this.tail.pre
		this.tail.next = nil
	} else {
		node.next.pre = node.pre
	}
	this.length--
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.store[key]
	if !exist {
		return -1
	}
	this.Delete(key)
	this.Put(node.key, node.value)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	this.Delete(key)
	if this.length+1 > this.capacity {
		this.Delete(this.tail.key)
	}
	node := LRUChainNode{key: key, value: value}
	if this.length == 0 {
		this.head = &node
		this.tail = &node
		this.store[key] = &node
		this.length++
		return
	}
	// 头部处理
	this.head.pre = &node
	node.next = this.head
	this.head = &node
	this.store[key] = &node
	this.length++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

