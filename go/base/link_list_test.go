package base

import (
	"fmt"
	"testing"
)

func display(ll *LinkList) {
	var fn = func(node *LinkListNode) {
		fmt.Printf("%v > ", node.value)
	}
	fmt.Print("link: ")
	ll.Visit(fn)
	fmt.Print("\n")
}

func TestLinkList(t *testing.T) {
	// a := CreateLinkList()
	// fmt.Println("a.length", a.length)
	// a.Insert(1)
	// display(a)
	// a.Insert(2)
	// a.Insert(3)
	// display(a)
	// // display(ra)
	// a.InsertAfter(23, a.Find(2))
	// a.InsertAfter(233, a.Find(2))
	// display(a)
	// a.DeleteByValue(233)
	// display(a)
	// a.Delete()
	// display(a)
	// a.Delete()
	// display(a)

	// b := CreateLinkList()
	// b.Insert(9)
	// b.Insert(6)
	// b.Insert(5)
	// b.Insert(2)

	// c := CreateLinkList()
	// c.Insert(8)
	// c.Insert(7)
	// c.Insert(4)
	// c.Insert(1)

	// d := MergeSortedLinkList(b, c)
	// display(d)

	e := CreateLinkList()
	for index := 10; index > 0; index-- {
		e.Insert(index)
	}
	display(e)
	f := ReverseLinkList(e.head)
	display(&LinkList{head: f, length: 10})
	ret, _ := GetRevKthFromLinkList(e.head, 0)
	println("ret=", ret)
}
