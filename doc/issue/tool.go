package issue

import "fmt"

type ListNode struct {
	Next *ListNode
	Data int
}

func (l *ListNode) Show() {
	if l == nil {
		fmt.Println("ListNode nil")
		return
	}
	var i int
	for {
		fmt.Printf("ShowList--no: %d, val: %d\n", i, l.Data)
		next := l.Next
		if next == nil {
			break
		}
		l = l.Next
		i++
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
