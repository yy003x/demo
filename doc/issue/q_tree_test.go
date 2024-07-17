package issue

import (
	"fmt"
	"testing"
)

func TestTraversal(t *testing.T) {
	root := generateTree(5, 1)
	printTree(root, 0)
	fmt.Printf("层序遍历: %v\n", LevelOrderTraversal(root))
	fmt.Printf("前序遍历: %v\n", PreOrderTraversal(root))
	fmt.Printf("中序遍历: %v\n", InOrderTraversal(root))
	fmt.Printf("后序遍历: %v\n", PostOrderTraversal(root))
}

// 层序遍历
func LevelOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	result := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Value)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return result
}

// 前序遍历
func PreOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	result := []int{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Value)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}

// 中序遍历
func InOrderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	result := []int{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Value)
		root = node.Right
	}
	return result
}

// 后序遍历
func PostOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	stack0 := []*TreeNode{}
	result := []int{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		stack0 = append(stack0, node)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	for len(stack0) > 0 {
		node := stack0[len(stack0)-1]
		stack0 = stack0[:len(stack0)-1]
		result = append(result, node.Value)
	}
	return result
}
func TestConstructTree(t *testing.T) {
	root := constructTree(3)
	printTree(root, 0)
	fmt.Println(preorderTraversal(root))
	fmt.Println(inorderTraversal(root))
	fmt.Println(postorderTraversal(root))
}

// 迭代生成二叉树
func constructTree(levels int) *TreeNode {
	if levels == 0 {
		return nil
	}
	value := 1
	root := &TreeNode{Value: value}
	queue := []*TreeNode{root}
	for level := 1; level < levels; level++ {
		lenQueue := len(queue)
		for i := 0; i < lenQueue; i++ {
			root := queue[0]
			queue = queue[1:]
			value++
			root.Left = &TreeNode{Value: value}
			value++
			root.Right = &TreeNode{Value: value}
			queue = append(queue, root.Left, root.Right)
		}
	}
	return root
}
