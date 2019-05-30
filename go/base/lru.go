package base

import (
	"errors"
)

// LruCache, 基于内存实现、不带过期时间
// 原理：map结构按照kv存储数据，双向链表保存数据新鲜度
// 扩展：支持过期时间可以增加一个双向链表按过期时间存储，
type LruCache struct {
	store    map[int]int
	list     *DualLinkList
	capacity int
	length   int
}

// NewLruCache return LRUCache ins
func NewLruCache(c int) *LruCache {
	return &LruCache{
		store:    map[int]int{},
		list:     &DualLinkList{length: 0, head: nil, tail: nil},
		capacity: c,
		length:   0,
	}
}

func (c *LruCache) Put(k int, v int) error {
	c.store[k] = v
	c.list.DeleteByValue(k)
	c.list.Append(k)
	if c.length >= c.capacity {
		node := c.list.Pop()
		delete(c.store, node.value)
	} else {
		c.length++
	}
	return nil
}

func (c *LruCache) Get(k int) (int, error) {
	if v, ok := c.store[k]; ok {
		c.list.DeleteByValue(k)
		c.list.Insert(k)
		return v, nil
	}
	return -1, errors.New("not found")
}

func (c *LruCache) Delete(k int) error {
	if _, ok := c.store[k]; ok {
		c.list.DeleteByValue(k)
		delete(c.store, k)
		c.length--
		return nil
	}
	return nil
}
