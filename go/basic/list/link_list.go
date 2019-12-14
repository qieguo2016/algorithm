package list

type LinkListNode struct {
	value int
	next  *LinkListNode
}

func (node *LinkListNode) IsEqual(target *LinkListNode) bool {
	return node.value == target.value && node.next == target.next
}

type LinkList struct {
	head   *LinkListNode
	length int
}

func CreateLinkList() *LinkList {
	return &LinkList{head: nil, length: 0}
}

// Insert 插入到头部
func (ll *LinkList) Insert(value int) *LinkListNode {
	node := LinkListNode{value: value}
	if ll.head != nil {
		node.next = ll.head
	}
	ll.head = &node
	ll.length++
	return &node
}

func (ll *LinkList) InsertAfter(value int, pos *LinkListNode) {
	node := LinkListNode{value: value}
	if pos.next != nil {
		node.next = pos.next
	}
	pos.next = &node
	ll.length++
}

func (ll *LinkList) Find(value int) *LinkListNode {
	for cur := ll.head; cur != nil; cur = cur.next {
		if value == cur.value {
			return cur
		}
	}
	return nil
}

// Delete 从头部删除
func (ll *LinkList) Delete() *LinkListNode {
	if ll.head == nil {
		return nil
	}
	node := ll.head
	ll.head = node.next
	ll.length--
	return node
}

// 根据取值删除
func (ll *LinkList) DeleteByValue(value int) *LinkListNode {
	cur := ll.head
	pre := ll.head
	for ; cur != nil; cur = cur.next {
		if value == cur.value {
			ll.length--
			if pre == ll.head {
				ll.head = nil
				return pre
			} else {
				pre.next = cur.next
				return cur
			}
		}
		pre = cur
	}
	return nil
}

// Visit 遍历高阶函数
func (ll *LinkList) Visit(fn func(node *LinkListNode)) {
	for cur := ll.head; cur != nil; cur = cur.next {
		fn(cur)
	}
}

type DualLinkList struct {
	head   *DualLinkListNode
	tail   *DualLinkListNode
	length int
}

type DualLinkListNode struct {
	pre   *DualLinkListNode
	next  *DualLinkListNode
	value int
}

func CreateDualLinkList() *DualLinkList {
	return &DualLinkList{head: nil, tail: nil, length: 0}
}

// Insert 插入到头部
func (dll *DualLinkList) Insert(value int) {
	node := DualLinkListNode{value: value}
	if dll.length == 0 {
		dll.head = &node
		dll.tail = &node
		dll.length++
		return
	}
	// 头部处理
	dll.head.pre = &node
	node.next = dll.head
	dll.head = &node
	dll.length++
}

// Append 追加到尾部
func (dll *DualLinkList) Append(value int) {
	node := DualLinkListNode{value: value}
	if dll.length == 0 {
		dll.tail = &node
		dll.head = &node
		dll.length++
		return
	}
	dll.tail.next = &node
	node.pre = dll.tail
	dll.tail = &node
	dll.length++
}

func (dll *DualLinkList) Find(value int) *DualLinkListNode {
	for cur := dll.head; cur != nil; cur = cur.next {
		if value == cur.value {
			return cur
		}
	}
	return nil
}

// Pop 从头部删除
func (dll *DualLinkList) Pop() *DualLinkListNode {
	if dll.length == 0 {
		return nil
	}
	if dll.length == 1 {
		node := dll.head
		dll.head = nil
		dll.tail = nil
		return node
	}
	node := dll.head
	dll.head = node.next
	dll.head.pre = nil
	dll.length--
	return node
}

// Shift 从尾部删除
func (dll *DualLinkList) Shift() *DualLinkListNode {
	if dll.length == 0 {
		return nil
	}
	if dll.length == 1 {
		node := dll.head
		dll.head = nil
		dll.tail = nil
		return node
	}
	node := dll.tail
	dll.tail = node.pre
	dll.tail.next = nil
	dll.length--
	return node
}

// DeleteByValue 根据取值删除
func (dll *DualLinkList) DeleteByValue(value int) *DualLinkListNode {
	for cur := dll.head; cur != nil; cur = cur.next {
		if value == cur.value {
			if cur.pre == nil {
				dll.head = dll.head.next
				if dll.head != nil {
					dll.head.pre = nil
				}
			} else {
				cur.pre.next = cur.next
			}
			if cur.next == nil {
				dll.tail = dll.tail.pre
				if dll.tail != nil {
					dll.tail.next = nil
				}
			} else {
				cur.next.pre = cur.pre
			}
			dll.length--
			return cur
		}
	}
	return nil
}

// Visit 遍历高阶函数
func (dll *DualLinkList) Visit(fn func(node *DualLinkListNode)) {
	for cur := dll.head; cur != nil; cur = cur.next {
		fn(cur)
	}
}

/****** go源码的双链表实现 ******/

// ListElement 链表节点，list指向链表
type ListElement struct {
	prev, next *ListElement
	list       *List
	Value      interface{}
}

// List 链表结构，首尾相连
type List struct {
	root ListElement
	len  int
}

// New 新建双链表
func New() *List {
	l := &List{
		root: ListElement{},
		len:  0,
	}
	l.root.prev = &l.root
	l.root.next = &l.root
	return l
}

// insert 将el插入到pos之后
func (l *List) insert(el, pos *ListElement) *ListElement {
	n := pos.next
	el.next = n
	n.prev = el
	pos.next = el
	el.prev = pos
	el.list = l
	l.len++
	return el
}

func (l *List) remove(el *ListElement) *ListElement {
	el.prev.next = el.next
	el.next.prev = el.prev
	el.prev = nil
	el.next = nil
	el.list = nil
	l.len--
	return el
}

// InsertBefore 插入到指定位置之前
func (l *List) InsertBefore(v interface{}, pos *ListElement) *ListElement {
	if pos.list != l {
		return nil
	}
	return l.insert(&ListElement{Value: v}, pos.prev)
}

// InsertAfter 插入到指定位置之前
func (l *List) InsertAfter(v interface{}, pos *ListElement) *ListElement {
	if pos.list != l {
		return nil
	}
	return l.insert(&ListElement{Value: v}, pos)
}

// Remove 删除节点
func (l *List) Remove(el *ListElement) interface{} {
	if el.list == l {
		l.remove(el)
	}
	return el.Value
}

// MoveBefore 移动节点到指定位置之前
func (l *List) MoveBefore(el, pos *ListElement) {
	if el.list != l || el == pos || pos.list != l {
		return
	}
	l.InsertBefore(l.Remove(el), pos)
}

// MoveAfter 移动节点到指定位置之后
func (l *List) MoveAfter(el, pos *ListElement) {
	if el.list != l || el == pos || pos.list != l {
		return
	}
	l.InsertAfter(l.Remove(el), pos)
}

// Len 长度
func (l *List) Len() int {
	return l.len
}

// Head 链表头节点
func (l *List) Head() *ListElement {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Tail 链表尾节点
func (l *List) Tail() *ListElement {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}
