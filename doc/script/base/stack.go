package base

import (
	"fmt"
	"sync"
)

type Item interface{}

type DefStackNode struct {
	Value Item
	Next  *DefStackNode
}
type DefStack struct {
	Length int
	Top    *DefStackNode
	Lock   *sync.RWMutex
}

func NewDefStack() *DefStack {
	return &DefStack{
		Length: 0,
		Top:    nil,
		Lock:   &sync.RWMutex{},
	}
}
func (s *DefStack) Show() {
	c := 0
	v := s.Top
	for v != nil {
		fmt.Println(v.Value)
		v = v.Next
		c++
	}
}

func (s *DefStack) Peek() Item {
	if s == nil || s.Length == 0 {
		return nil
	}
	return s.Top.Value
}

func (s *DefStack) Push(v Item) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.Top = &DefStackNode{v, s.Top}
	s.Length++
}

func (s *DefStack) Pop() Item {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if s.Length == 0 {
		return nil
	}
	n := s.Top
	s.Top = n.Next
	s.Length--
	return n.Value
}

type ItemStack struct {
	Items []Item
	Lock  *sync.RWMutex
}

func NewItemStack() *ItemStack {
	return &ItemStack{
		Items: []Item{},
		Lock:  &sync.RWMutex{},
	}
}

func (s *ItemStack) Print() {
	fmt.Printf("%+v \n", s.Items)
}

func (s *ItemStack) Push(d Item) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.Items = append(s.Items, d)
}

func (s *ItemStack) Pop() Item {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if len(s.Items) == 0 {
		return nil
	}
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[0 : len(s.Items)-1]
	return item
}
