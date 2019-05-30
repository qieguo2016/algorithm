package base

import (
	"sort"
)

// SmallRootDown 小根堆向下调整
func SmallRootDown(target []int, from int, to int) {
	parent := from
	left := 2*from + 1
	if left >= to || left < 0 { // < 0: int overflow
		return
	}

	// 获取左右节点的较小值
	smaller := left
	if right := left + 1; right < to && target[left] > target[right] {
		smaller = right
	}
	// 比较父节点与较小值，若父节点小则符合小根堆要求
	if target[parent] < target[smaller] {
		return
	}

	// 父节点取小节点
	swap(target, smaller, parent)
	SmallRootDown(target, parent, to)
}

// NewSmallRootHeap 构建小根堆
func NewSmallRootHeap(target []int) {
	l := len(target)
	// 从各层根节点向下调整
	for i := l/2 - 1; i >= 0; i-- {
		SmallRootDown(target, i, l)
	}
}

/********** 以下是源码的实现方式 *********/

// HeapInterface
type HeapInterface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

// 向下调整各层的根堆
func down(h HeapInterface, from int, to int) {
	parent := from
	for {
		left := 2*parent + 1
		if left >= to || left < 0 { // < 0: int overflow
			break
		}
		smaller := left
		// 比较左右叶子节点，取小
		if right := left + 1; right < to && h.Less(right, left) {
			smaller = right
		}
		// 比较父节点与小节点，如果父节点小于小节点则已经符合要求，跳出循环
		if h.Less(parent, smaller) {
			break
		}
		// 父节点改为小节点，根堆结构改变，此时才需要继续调整下级根堆
		h.Swap(parent, smaller)
		parent = smaller
	}
}

// 向上调整各层的根堆
func up(h HeapInterface, from int) {
	for {
		parent := (from - 1) / 2
		// from 大于 parent，符合小根堆定义，无需调整
		if parent == from || !h.Less(from, parent) {
			break
		}
		h.Swap(parent, from)
		from = parent
	}
}

// NewHeap 大小根堆区别在h.Less()实现上
func NewHeap(h HeapInterface) {
	length := h.Len()
	for i := length/2 - 1; i >= 0; i-- {
		down(h, i, length)
	}
}
