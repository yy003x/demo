package case_test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	nums := []int{-1, 1, 3, -3, -5, 0, 4, 6}
	fmt.Print(MaxArr(nums))
}

//  clent-log
//   segement  user - log - trace - sync server-log  task - hive/T+1  hbase/Seconds / ES index rpcid
//  A 40% B 40% C20%

//   segement
//  A 100%  - action  A - 10
// B 100%  - action B    - 10
// C 100%  - action C  -10

//   segement
//  A 100%  - action  A
// B 50% C 50% - action B/C
// D 100%  - action D

// DMP  tracing  链路追踪 case bug
// id ROI   100w GMV
func MaxArr(nums []int) int {
	ans := 0
	n := len(nums)
	left := 0
	for right := left; right < n; right++ {
		if nums[left] <= 0 {
			left++
		}
		ans += nums[right]
	}
	return ans
}

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func CheckTree(root *TreeNode) bool {
	leftHigh := High(root.Left, 1)
	rightHigh := High(root.Right, 1)
	if rightHigh-leftHigh > 1 || leftHigh-rightHigh > 1 {
		return false
	}
	return true
}

func High(node *TreeNode, high int) int {
	lhigh, rhigh := 0, 0
	high = high + 1
	if node.Left != nil {
		lhigh = High(node.Left, high)
	}
	if node.Right != nil {
		rhigh = High(node.Left, high)
	}
	return max(lhigh, rhigh)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
