package cache

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

// NewLFUCache constructor
func NewLFUCache(capacity int) LFUCache {
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
