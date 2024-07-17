package issue

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestGetIntersectionNode(t *testing.T) {
	// 创建相交部分
	intersection := &ListNode{Val: 6}
	intersection.Next = &ListNode{Val: 7}
	// 创建链表A
	headA := &ListNode{Val: 1}
	headA.Next = &ListNode{Val: 2}
	headA.Next.Next = intersection
	// 创建链表B
	headB := &ListNode{Val: 3}
	headB.Next = &ListNode{Val: 4}
	headB.Next.Next = &ListNode{Val: 5}
	headB.Next.Next.Next = intersection
	// 打印链表
	fmt.Print("List A: ")
	PrintList(headA)
	fmt.Print("List B: ")
	PrintList(headB)
	// 找到相交节点
	intersectionNode := GetIntersectionNode(headA, headB)
	PrintList(intersectionNode)
}

// 相交链表
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

// 环形链表
func HasCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	// 没有环
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// 打印链表
func PrintList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}
