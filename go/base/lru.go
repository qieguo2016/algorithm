package base

// LruCache 基于内存实现、不带过期时间
// 原理：map结构按照kv存储数据，双向链表保存数据新鲜度
// 扩展：支持过期时间可以增加一个双向链表按过期时间存储，

// LRUChainNode 链表节点
type LRUChainNode struct {
	pre   *LRUChainNode
	next  *LRUChainNode
	key   int
	value int
	ts    int32
}

// LRUCache 结构
type LRUCache struct {
	capacity int
	length   int
	store    map[int]*LRUChainNode
	head     *LRUChainNode
	tail     *LRUChainNode
}

// NewLRUCache constructor
func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		length:   0,
		store:    map[int]*LRUChainNode{},
	}
}

// Delete 删除key
func (c *LRUCache) Delete(key int) {
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
	if node.pre == nil {
		c.head = c.head.next
		c.head.pre = nil
	} else {
		node.pre.next = node.next
	}
	if node.next == nil {
		c.tail = c.tail.pre
		c.tail.next = nil
	} else {
		node.next.pre = node.pre
	}
	c.length--
}

// Get 获取kv
func (c *LRUCache) Get(key int) int {
	node, exist := c.store[key]
	if !exist {
		return -1
	}
	c.Delete(key)
	c.Put(node.key, node.value)
	return node.value
}

// Put 插入
func (c *LRUCache) Put(key int, value int) {
	c.Delete(key)
	if c.length+1 > c.capacity {
		c.Delete(c.tail.key)
	}
	node := LRUChainNode{key: key, value: value}
	if c.length == 0 {
		c.head = &node
		c.tail = &node
		c.store[key] = &node
		c.length++
		return
	}
	// 头部处理
	c.head.pre = &node
	node.next = c.head
	c.head = &node
	c.store[key] = &node
	c.length++
}
