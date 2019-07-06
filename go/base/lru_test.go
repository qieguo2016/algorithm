package base

import (
	"fmt"
	"testing"
)

func TestLru(t *testing.T) {
	cache := NewLRUCache(1)
	cache.Put(1, 11)
	cache.Put(2, 22)
	cache.Put(3, 33)
	fmt.Println("put success")

	v := cache.Get(2)
	fmt.Printf("get (2) return %v\n", v) // 返回  22

	v = cache.Get(2)
	fmt.Printf("get (2) return %v\n", v) // 返回  22

	cache.Put(4, 44) // 该操作会使(2,22) 作废
	fmt.Println("put 4 success")

	v = cache.Get(2)
	fmt.Printf("get (2) return %v \n", v) // 返回 -1 (未找到)

	v = cache.Get(3)
	fmt.Printf("get (3) return %v\n", v) // 返回  33
}
