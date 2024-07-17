package issue

import (
	"fmt"
	"testing"
)

// TreeNode 定义二叉树节点结构体
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func TestGenerateTree(t *testing.T) {
	root := generateTree(4, 1)
	printTree(root, 0)
	// depth := maxDepth(root)
	// fmt.Println(depth)
	invertTree(root)
	printTree(root, 0)
	// fmt.Println(preorderTraversal(root))
	// fmt.Println(inorderTraversal(root))
	// fmt.Println(postorderTraversal(root))
}

func TestTreeSort(t *testing.T) {
	nums := []int{9, 5, 3, 8, 2, 7, 6, 4, 1}
	fmt.Println(nums)
	sarr := TreeSort(nums)
	fmt.Println(sarr)
	bst := SortedArrayToBST(sarr)
	fmt.Println(inorderTraversal(bst))
	printTree(bst, 0)
}

// 将有序数组转换为二叉搜索树
func SortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}
func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (right-left)/2 + left
	root := &TreeNode{Value: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}

// 二叉树排序
func TreeSort(nums []int) []int {
	var root *TreeNode
	for _, num := range nums {
		root = insert(root, num)
	}
	return inorderTraversal(root)
}
func insert(root *TreeNode, num int) *TreeNode {
	if root == nil {
		return &TreeNode{Value: num}
	}
	if num < root.Value {
		root.Left = insert(root.Left, num)
	} else {
		root.Right = insert(root.Right, num)
	}
	return root
}

func TestIsSymmetric(t *testing.T) {
	root := &TreeNode{Value: 1}
	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: 2}
	root.Left.Left = &TreeNode{Value: 3}
	root.Left.Right = &TreeNode{Value: 4}
	root.Right.Left = &TreeNode{Value: 4}
	root.Right.Right = &TreeNode{Value: 3}
	printTree(root, 0)
	fmt.Println(isSymmetric(root))
}

// 对称二叉树
func isSymmetric(root *TreeNode) bool {
	var check func(l, r *TreeNode) bool
	check = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l == nil || r == nil {
			return false
		}
		if l.Value == r.Value && check(l.Left, r.Right) && check(l.Right, r.Left) {
			return true
		}
		return false
	}
	return check(root, root)
}

// 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Right)
	right := invertTree(root.Left)
	root.Right = right
	root.Left = left
	return root
}

// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	return depth
}

// 生成二叉树
func generateTree(level, currentValue int) *TreeNode {
	if level == 0 {
		return nil
	}
	root := &TreeNode{Value: currentValue}
	root.Left = generateTree(level-1, currentValue*2)
	root.Right = generateTree(level-1, currentValue*2+1)
	return root
}

//前序遍历
func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node != nil {
			result = append(result, node.Value)
			preorder(node.Left)
			preorder(node.Right)
		}
	}
	preorder(root)
	return result
}

//中序遍历
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)
			result = append(result, node.Value)
			inorder(node.Right)
		}
	}
	inorder(root)
	return result
}

//后序遍历
func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node != nil {
			postorder(node.Left)
			postorder(node.Right)
			result = append(result, node.Value)
		}
	}
	postorder(root)
	return result
}

// 打印二叉树
func printTree(root *TreeNode, level int) {
	if root == nil {
		return
	}
	printTree(root.Right, level+1)
	fmt.Printf("%*s%d\n", level*4, "", root.Value)
	printTree(root.Left, level+1)
}
