package base

import (
	"container/heap"
	"fmt"
	"testing"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestHeap(t *testing.T) {
	h := &MaxHeap{2, 1, 5, 4, 3}
	heap.Init(h)

	heap.Push(h, 6)
	fmt.Printf("堆顶元素：%d\n", (*h)[0])

	fmt.Println(h)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
