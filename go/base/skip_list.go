package base

import (
	"math/rand"
)

type SkipList struct {
	MaxLevel int // 现有层级
	Length   int
	Head     *SkipListNode

	levelLimit int
	p          float64
}

type SkipListNode struct {
	Index int
	Data  interface{}
	Pre   *SkipListNode
	Level []*SkipListLevel
}

type SkipListLevel struct {
	Next *SkipListNode
	Span int
}

func NewSkipList() *SkipList {
	levelLimit := 10
	p := 0.25
	s := &SkipList{
		MaxLevel:   0, // level: [0, n)
		Length:     0,
		Head:       &SkipListNode{Level: make([]*SkipListLevel, levelLimit)},
		levelLimit: levelLimit,
		p:          p,
	}
	for i := 0; i < levelLimit; i++ {
		s.Head.Level[i] = &SkipListLevel{}
	}
	return s
}

func (s *SkipList) randomLevel() int {
	level := 0
	for rand.Float64() < s.p && level < s.levelLimit-1 {
		level++
	}
	return level // [0, n)
}

func (s *SkipList) Insert(index int, data interface{}) {
	n := s.randomLevel() // 随机层级
	if n > s.MaxLevel {
		s.MaxLevel = n
	}
	node := &SkipListNode{
		Index: index,
		Data:  data,
		Level: make([]*SkipListLevel, n+1),
	}
	preNodeList := make([]*SkipListNode, n+1)
	spanList := make([]int, n+1)
	preNode := s.Head
	// 先遍历寻找目标位，如果目标已经存在，则直接替换data即可
	for n >= 0 {
		span := 0
		next := preNode.Level[n].Next
		for next != nil { // 为了span叠加计数，index比较条件分开
			if next.Index > index {
				break
			}
			if next.Index == index { // 找到目标位、直接修改data
				next.Data = data
				return
			}
			span += preNode.Level[n].Span
			preNode = next
			next = next.Level[n].Next
		}
		preNodeList[n] = preNode // 存储各层的前节点
		spanList[n] = span       // 存储各层的距离值
		n--
	}
	for i := 0; i < len(preNodeList); i++ {
		// 先处理前后关系
		node.Level[i] = &SkipListLevel{Next: preNodeList[i].Level[i].Next}
		preNodeList[i].Level[i].Next = node
		if i <= 0 {
			// 底层双向链表连接
			node.Pre = preNodeList[0]
			if node.Level[i].Next != nil {
				node.Level[i].Next.Pre = node
				node.Level[i].Span = 1 // 底层直接赋1
			}
			preNodeList[i].Level[i].Span = 1
		} else {
			// 由下层关系递推得到
			span := preNodeList[i-1].Level[i-1].Span + spanList[i-1]
			if node.Level[i].Next != nil {
				node.Level[i].Span = preNodeList[i].Level[i].Span - span + 1
			}
			preNodeList[i].Level[i].Span = span
		}
	}
	s.Length++
}

func (s *SkipList) Delete(index int) interface{} {
	preNode := s.Head
	n := s.MaxLevel // 从最高层往下找
	preNodeList := make([]*SkipListNode, n+1)
	var node *SkipListNode
	for n >= 0 {
		next := preNode.Level[n].Next
		for next != nil {
			if next.Index >= index {
				break
			}
			preNode = next
			next = next.Level[n].Next
		}
		if node == nil && next != nil && next.Index == index {
			node = next
		}
		preNodeList[n] = preNode // 存储各层的前节点
		n--
	}
	if node == nil {
		return nil
	}
	pre := node.Pre
	pre.Level[0].Next = node.Level[0].Next
	pre.Level[0].Next.Pre = pre
	node.Level[0].Next = nil
	node.Pre = nil
	for i := 1; i < len(node.Level); i++ {
		preNodeList[i].Level[i].Next = node.Level[i].Next
		preNodeList[i].Level[i].Span += node.Level[i].Span - 1
	}
	return node.Data
}

func (s *SkipList) searchNode(index int) *SkipListNode {
	preNode := s.Head
	n := s.MaxLevel // 从最高层往下找
	for n >= 0 {
		next := preNode.Level[n].Next
		for next != nil {
			if next.Index == index {
				return next
			}
			if next.Index > index {
				break
			}
			preNode = next
			next = next.Level[n].Next
		}
		n--
	}
	return nil
}

func (s *SkipList) Search(index int) interface{} {
	node := s.searchNode(index)
	if node == nil {
		return nil
	}
	return node.Data
}

func (s *SkipList) Range(startIndex int, endIndex int) []interface{} {
	ret := []interface{}{}
	startNode := s.searchNode(startIndex)
	if startNode == nil {
		return ret
	}
	ret = append(ret, startNode.Data)
	for next := startNode.Level[0].Next; next != nil && next.Index <= endIndex; next = next.Level[0].Next {
		ret = append(ret, next.Data)
	}
	return ret
}

func (s *SkipList) rankNode(rank int, delta int) []*SkipListNode {
	ret := []*SkipListNode{}
	preNode := s.Head
	n := s.MaxLevel // 从最高层往下找
	start := 0
	var node *SkipListNode
	for n >= 0 {
		for preNode.Level[n] != nil && start+preNode.Level[n].Span <= rank {
			start += preNode.Level[n].Span
			if preNode.Level[n].Next == nil {
				break
			}
			preNode = preNode.Level[n].Next
		}
		if start == rank {
			node = preNode
			break
		}
		n--
		continue
	}
	if node == nil {
		return ret
	}
	ret = append(ret, node)
	for delta > 0 {
		node = node.Level[0].Next
		if node != nil {
			ret = append(ret, node)
		}
		delta--
	}
	for delta < 0 {
		node = node.Pre
		if node != nil {
			ret = append(ret, node)
		}
		delta--
	}
	return ret
}

func (s *SkipList) Rank(rank int, delta int) []interface{} {
	nodes := s.rankNode(rank, delta)
	ret := []interface{}{}
	for _, el := range nodes {
		ret = append(ret, el.Data)
	}
	return ret
}
