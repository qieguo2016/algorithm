package base

import (
	"fmt"
	"testing"
)

func TestLru(t *testing.T) {
	cache := NewLruCache(1)
	cache.Put(2, 22)
	fmt.Println("put (2, 22) success")

	v, err := cache.Get(2)
	fmt.Printf("get (2) return %v, err=%v\n", v, err) // 返回  22

	cache.Put(3, 33) // 该操作会使(2,22) 作废
	fmt.Println("put (3, 33)  success")

	v, err = cache.Get(2)
	fmt.Printf("get (2) return %v, err=%v\n", v, err) // 返回 -1 (未找到)

	v, err = cache.Get(3)
	fmt.Printf("get (3) return %v, err=%v\n", v, err) // 返回  33
}
