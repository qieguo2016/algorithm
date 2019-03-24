package base

type DualLinkList struct {
	Head  *DualLinkList
	Tail  *DualLinkList
	Count int
}

type DualLinkListNode struct {
	Pre   *DualLinkList
	Next  *DualLinkList
	Value int
}

func (dll *DualLinkList) Insert(value int) {
	node := DualLinkListNode{Value: value}
	if dll.Tail != nil {
		dll.Tail.Next = &node
		node.Pre = dll.Tail
	}
	dll.Tail = &node
	if dll.Head != nil {
		dll.Head = &node
	}
	dll.Count++
}

func (this *DualLinkList) InsertAfter(value int, pos *DualLinkListNode) {

}

func (this *DualLinkList) Delete(node *DualLinkListNode) {
	if node.Pre != nil {
		node.Pre.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Pre = node.Pre
	}
	if this.Head == node {
		this.Head = nil
	}
	if this.Tail == node {
		this.Tail = node
	}
	this.Count--
}

func (this *DualLinkList) Search(value int) {

}

func (this *DualLinkList) Visit() {

}

func (this *DualLinkList) VisitFromTail() {

}
