package concurrency

import (
	"sync"
	"sync/atomic"
)

// Once sync底层使用锁和原子自增实现once效果
type Once struct {
	mu    sync.Mutex
	count int32
}

// Do 提供一个do方法
func (o *Once) Do(f func()) {
	// 使用cas限制第一次也OK？
	// 不ok，因为两个g同时进来的时候，其中一个没有执行f，另外一个执行了f，两个g后面取到的值是有差异的，一致性不满足
	// 而采用源码的方式，在抢互斥锁的时候另外一个线程堵塞，保证后续运行时都是在f执行后的状态
	// if atomic.CompareAndSwapInt32(&o.count, 0, 1) {
	// 	f()
	// }

	// 源码实现，先原子取数，然后再抢锁，抢到之后判断是否执行过f，若未执行则调用f，然后再原子+1
	if atomic.LoadInt32(&o.count) == 1 {
		return
	}
	o.mu.Lock()
	if o.count == 0 {
		f()
		atomic.StoreInt32(&o.count, 1)
	}
	o.mu.Unlock()
}
