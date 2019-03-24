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
	pre  *LRUChainNode
	next *LRUChainNode
	key  int
	ts   int32
}

type CacheObject struct {
	value int
	node  *LRUChainNode
}

type LRUCache struct {
	cap   int
	num   int
	store map[int]*CacheObject
	head  *LRUChainNode
	tail  *LRUChainNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		num:   0,
		store: map[int]*CacheObject{},
	}
}

func (this *LRUCache) updateLRUChain(node *LRUChainNode, isPut bool) {
	// 从原位置删掉
	if node.pre != nil {
		node.pre.next = node.next
	}
	if node.next != nil {
		node.next.pre = node.pre
	}
	// 处理队尾
	if this.tail == node && node.pre != nil {
		this.tail = node.pre
	}
	// 插入到head
	node.pre = nil
	node.next = this.head
	if this.head != nil {
		this.head.pre = node
	}
	this.head = node
	if !isPut {
		return
	}
	if this.tail == nil {
		this.tail = node
		this.num++
		return
	}
	if this.num+1 > this.cap {
		delete(this.store, this.tail.key)
		this.tail = this.tail.pre
		return
	}
	this.num++
}

func (this *LRUCache) Get(key int) int {
	obj, exist := this.store[key]
	if !exist {
		return -1
	}
	this.updateLRUChain(obj.node, false)
	return obj.value
}

func (this *LRUCache) Put(key int, value int) {
	obj, exist := this.store[key]
	if exist {
		obj.value = value
		this.updateLRUChain(obj.node, false)
		return
	}
	obj = &CacheObject{value: value}
	obj.node = &LRUChainNode{key: key}
	this.store[key] = obj
	this.updateLRUChain(obj.node, true)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

