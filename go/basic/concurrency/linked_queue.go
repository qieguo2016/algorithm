/*
	lock free linked queue.
	ref:
		1. http://ddrv.cn/a/591069
		2. https://coolshell.cn/articles/8239.html
*/

package concurrent

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type LinkedQueueNode struct {
	Value interface{}
	Next  *LinkedQueueNode
}

func (node *LinkedQueueNode) casNext(oldV, newV *LinkedQueueNode) bool {
	return atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&node.Next)),
		unsafe.Pointer(oldV),
		unsafe.Pointer(newV),
	)
}

func (node *LinkedQueueNode) loadNext() *LinkedQueueNode {
	return (*LinkedQueueNode)(atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&node.Next)),
	))
}

type LinkedQueue struct {
	head *LinkedQueueNode
	tail *LinkedQueueNode
	size int64
	m    sync.Mutex
}

func NewLinkedQueue() *LinkedQueue {
	dummy := &LinkedQueueNode{}
	dummy.Value = nil
	dummy.Next = nil
	return &LinkedQueue{ // like container/list, use same node
		head: dummy,
		tail: dummy,
	}
}

func (queue *LinkedQueue) casTail(oldV, newV *LinkedQueueNode) bool {
	return atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&queue.tail)),
		unsafe.Pointer(oldV),
		unsafe.Pointer(newV),
	)
}

func (queue *LinkedQueue) casHead(oldV, newV *LinkedQueueNode) bool {
	return atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&queue.head)),
		unsafe.Pointer(oldV),
		unsafe.Pointer(newV),
	)
}

func (queue *LinkedQueue) loadHead() *LinkedQueueNode {
	return (*LinkedQueueNode)(atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&queue.head)),
	))
}

func (queue *LinkedQueue) loadTail() *LinkedQueueNode {
	return (*LinkedQueueNode)(atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&queue.tail)),
	))
}

func (queue *LinkedQueue) Enqueue(v interface{}) bool {
	newNode := &LinkedQueueNode{Value: v, Next: nil}
	var tail, next *LinkedQueueNode
	for {
		// use atomic load and cas
		tail = queue.loadTail()
		next = tail.loadNext()
		if tail == queue.loadTail() { // double check
			if next == nil { // queue tail
				if tail.casNext(next, newNode) { // link to queue
					break
				}
			} else {
				queue.casTail(tail, next) // move tail pointer to real tail
			}
		}
	}

	queue.casTail(tail, newNode) // failure is ok, another thread has update
	atomic.AddInt64(&queue.size, 1)
	return true
}

func (queue *LinkedQueue) Dequeue() interface{} {
	var head, tail, first *LinkedQueueNode
	for {
		// use atomic load and cas
		head = queue.loadHead()  // dummy
		tail = queue.loadTail()  // dummy
		first = head.loadNext()  // nil
		if head == queue.loadHead() { // double check
			if first == nil { // empty list
				return nil
			}
			if head == tail { // empty list
				queue.casTail(tail, first) // move tail to real pointer
				continue
			}
			if queue.casHead(head, first) {
				break
			}
		}
	}

	atomic.AddInt64(&queue.size, -1)
	return first.Value
}

func (queue *LinkedQueue) Size() int64 {
	return atomic.LoadInt64(&queue.size)
}

func (queue *LinkedQueue) EnqueueWithLock(v interface{}) bool {
	newNode := &LinkedQueueNode{Value: v, Next: nil}
	queue.m.Lock()
	defer queue.m.Unlock()
	tail := queue.tail
	tail.Next = newNode
	queue.tail = newNode
	queue.size += 1
	return true
}

func (queue *LinkedQueue) DequeueWithLock() interface{} {
	var head, tail, first *LinkedQueueNode
	queue.m.Lock()
	defer queue.m.Unlock()
	head = queue.head
	tail = queue.tail
	first = head.Next
	if head == tail {
		return nil
	}
	queue.head = first
	head.Next = nil
	queue.size -= 1
	return first.Value
}
