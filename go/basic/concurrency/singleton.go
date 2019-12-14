package concurrency

import (
	"sync"
	"time"
)

// Singleton 单例
type Singleton struct {
	Value int64
}

var s *Singleton

var _ = new(sync.Once)
var once = new(Once)

// GetGlobalSingle 使用once来实现全局单例
func GetGlobalSingle() *Singleton {
	once.Do(func() {
		s = &Singleton{
			Value: time.Now().Unix(),
		}
	})
	return s
}
