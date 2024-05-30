package base

import "fmt"

type SimpleList struct {
	data string
	next *SimpleList
}

func TestList() {
	l := NewSimpleList("aaa")
	fmt.Println(l)
	len := l.Length()
	fmt.Println(len)
	l = l.InsertNode("bbb", 0)
	l = l.InsertNode("ccc", 20)
	l = l.InsertNode("ddd", 4)
	l = l.InsertNode("eee", 0)
	a := l.FindIndex("ddd")
	fmt.Println("ddd position:", a)
	b := l.FindNode(0)
	fmt.Println("0 node:", b)
	fmt.Println(l)
	l.PrintAll()
}

func NewSimpleList(val string) *SimpleList {
	return &SimpleList{data: val, next: nil}
}

func (l *SimpleList) Length() int {
	len := 0
	for l.next != nil {
		len++
		l = l.next
	}
	if len == 0 && l.data != "" {
		len++
	}
	return len
}
func (l *SimpleList) FindIndex(val string) int {
	idx := 0
	for {
		if l.next == nil || val == l.data {
			break
		}
		l = l.next
		idx++
	}
	return idx
}

func (l *SimpleList) FindNode(idx int) *SimpleList {
	len := l.Length()
	i := 0
	if idx >= len || idx < i {
		return nil
	}
	for {
		if i == idx {
			break
		}
		l = l.next
		i++
	}
	return l
}

func (l *SimpleList) InsertNode(val string, idx int) *SimpleList {
	len := l.Length()
	nl := NewSimpleList(val)
	// nil list
	if len == 0 {
		return nl
	}
	//head
	if idx <= 0 {
		nl.next = l
		return nl
	} else {
		sn := l.FindNode(idx)
		if sn == nil {
			sn = l.FindNode(len - 1)
		}
		//tail
		if idx >= len {
			sn.next = nl
			return l
		}
		tmpNode := sn.next
		sn.next = nl
		nl.next = tmpNode
		return l
	}
}

func (l *SimpleList) PrintAll() {
	for l.next != nil {
		fmt.Println(l.data)
		l = l.next
	}
}

func (l *SimpleList) DeleteNode(idx int) *SimpleList {
	len := l.Length()
	if idx < 0 || idx >= len {
		return nil
	}
	for i := 0; i < idx-1; i++ {
		l = l.next
	}
	return l
}
