package base

import (
	"fmt"
)

func Testxs() {
	l := NewLinked()
	l.Insert(0, 111)
	l.Insert(1, "123")
	l.Show()
	l.Insert(1, "234")
	l.Insert(10, "10")
	l.Insert(2, "101")
	l.Insert(3, "101")
	l.Insert(4, "101")
	l.Insert(5, "101")

	// l.Index(4)
	// l.Show()
	// l.Delete(3)
	// l.Delete(0)
	l.Show()
	// l.Reverse()
}

type Data interface{}

type Node struct {
	Data Data
	Next *Node
}

type Linked struct {
	Size int
	Head *Node
	Tail *Node
}

func NewNode(d Data) *Node {
	return &Node{
		Data: d,
	}
}

func NewLinked() *Linked {
	return &Linked{
		Size: 0,
		Head: nil,
		Tail: nil,
	}
}

func (l *Linked) Show() {
	c := 0
	n := l.Head
	for n != nil {
		fmt.Printf("idx: %d, val: %+v \n", c, n.Data)
		n = n.Next
		c++
	}
	fmt.Printf("linked: %+v, size: %d \n", l, c)
}

func (l *Linked) Index(idx int) *Node {
	if idx < 0 || idx >= l.Size {
		return nil
	}
	cur := l.Head
	for i := 0; i < idx; i++ {
		if cur == nil {
			break
		}
		cur = cur.Next
	}
	fmt.Printf("index: %d, val: %+v \n", idx, cur)
	return cur
}

func (l *Linked) Search(data Data) int {
	return 0
}

func (l *Linked) Append(data Data) *Linked {
	return l
}

func (l *Linked) Insert(index int, data Data) *Linked {
	newNode := NewNode(data)
	curNode := l.Index(index - 1)
	if curNode == nil {
		if l.Size == 0 {
			l.Head = newNode
			l.Tail = newNode
		} else {
			if index == 0 {
				newNode.Next = l.Head
				l.Head = newNode
			} else {
				l.Tail.Next = newNode
				l.Tail = newNode
			}
		}
	} else {
		newNode.Next = curNode.Next
		curNode.Next = newNode
	}
	l.Size++
	fmt.Printf("insert: %d, link: %+v \n", index, l)
	return l
}

func (l *Linked) Delete(index int) *Linked {
	prevNode := l.Index(index - 1)
	if prevNode == nil || prevNode.Next == nil {
		return l
	}
	if index == 0 {
		l.Head = l.Head.Next
		l.Size--
	}
	curNode := prevNode.Next
	prevNode.Next = curNode.Next
	if prevNode.Next == nil {
		l.Tail = prevNode
	}
	l.Size--
	fmt.Printf("delete: %d, val: %+v \n", index, curNode)
	return l
}

func (l *Linked) Reverse() *Linked {
	head := l.Head
	if head == nil {
		return l
	}
	for head != nil {
		head = head.Next
		fmt.Printf("%+v", head)
	}
	return l
}
