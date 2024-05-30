package issue

import "fmt"

/* @Title 两数相加

给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

*/

func NumInit(num []int) *ListNode {
	var l *ListNode
	for i := len(num) - 1; i >= 0; i-- {
		node := &ListNode{
			Data: num[i],
		}
		if l == nil {
			l = node
			continue
		}
		node.Next = l
		l = node
	}
	return l
}

func Q3() {
	num1 := []int{2, 3, 4, 5}
	num2 := []int{9, 8, 7, 6}
	fmt.Println(num1)
	l1 := NumInit(num1)
	l2 := NumInit(num2)
	l1.Show()
	l2.Show()
}

func TwoNumAdd(l1, l2 *ListNode) *ListNode {
	var node, res *ListNode
	var bit int
	for l1 != nil || l2 != nil {
		var n1, n2 int
		if l1 != nil {
			n1 = l1.Data
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Data
			l2 = l2.Next
		}
		sum := n1 + n2 + bit
		bit = sum / 10
		no := sum % 10
		if res == nil {
			res = &ListNode{
				Data: no,
			}
		}
	}
	return node
}
