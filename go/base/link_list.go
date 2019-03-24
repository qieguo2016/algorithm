package base

type LinkListNode struct {
	value int
	next  *LinkListNode
}

type LinkList struct {
	head  *LinkListNode
	count int
}

func CreateLinkList() *LinkList {
	return &LinkList{head: nil, count: 0}
}

// 插入到头部
func (ll *LinkList) Insert(value int) *LinkListNode{
	node := LinkListNode{value: value}
	if ll.head != nil {
		node.next = ll.head
	}
	ll.head = &node
	ll.count++
	return &node
}

func (ll *LinkList) InsertAfter(value int, pos *LinkListNode) {
	node := LinkListNode{value: value}
	if pos.next != nil {
		node.next = pos.next
	}
	pos.next = &node
	ll.count++
}

func (ll *LinkList) Find(value int) *LinkListNode{
	for cur := ll.head; cur != nil; cur = cur.next {
		if value == cur.value {
			return &cur
		}
	}
	return nil
}

// 从头部删除
func (ll *LinkList) Delete() *LinkListNode{
	if ll.head == nil {
		return nil
	}
	node := ll.head
	ll.head = node.next
	ll.count--
	return node
}

// 根据取值删除
func (ll *LinkList) DeleteByValue(value int) *LinkListNode {
	for cur := ll.head, pre:=ll.head ; cur != nil; cur = cur.next {
		if value == cur.value {
			ll.count--
			if pre == ll.head {
				ll.head = nil
				return pre
			} else {
				pre.next = cur.next
				return cur
			}
		}
	}
	return nil
}

func (ll *LinkList) Visit(fn func(node *LinkListNode))  {
	for cur := ll.head; cur != nil; cur = cur.next {
		fn(cur)
	}
}