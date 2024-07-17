package issue

import (
	"fmt"
	"testing"
)

func Test206(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Print("Original list: ")
	PrintList(head)

	reversedHead := reverseListRecursive(head)

	fmt.Print("Reversed list: ")
	PrintList(reversedHead)
}

// 反转链表
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		// 暂存当前节点的下一个节点指针
		tmpNext := curr.Next
		// 反转当前节点的指针到上一个节点
		curr.Next = prev
		// 移动 prev 指针到当前节点
		prev = curr
		// 移动 current 指针到下一个节点
		curr = tmpNext
	}
	return prev
}

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
